package bench

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"
	"slices"
	"sort"
	"strings"
	"sync/atomic"
	"time"

	"github.com/urfave/cli/v3"
	v1 "github.com/vinceanalytics/vince/gen/proto/go/vince/v1"
	"github.com/vinceanalytics/vince/internal/cmd/ansi"
	"github.com/vinceanalytics/vince/internal/cmd/auth"
	"github.com/vinceanalytics/vince/internal/entry"
	"github.com/vinceanalytics/vince/internal/geoip"
	"github.com/vinceanalytics/vince/internal/klient"
	"github.com/vinceanalytics/vince/internal/referrer"
	"github.com/vinceanalytics/vince/internal/ua"
	"golang.org/x/sync/errgroup"
)

func CMD() *cli.Command {
	return &cli.Command{
		Name:  "bench",
		Usage: "Load test a vince instance",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "users,u",
				Usage:   "Number of concurrent users",
				Value:   1,
				EnvVars: []string{"VINCE_BENCH_USERS"},
			},
			&cli.StringFlag{
				Name:    "event,e",
				Usage:   "Name of the event to send",
				Value:   "pageview",
				EnvVars: []string{"VINCE_BENCH_EVENT"},
			},
			&cli.StringSliceFlag{
				Name:    "paths",
				Usage:   "Url path visited by the users",
				Value:   []string{"/", "/pricing", "/company"},
				EnvVars: []string{"VINCE_BENCH_PATHS"},
			},
			&cli.DurationFlag{
				Name:    "duration,d",
				Usage:   "How long to run the benchmark",
				Value:   time.Second,
				EnvVars: []string{"VINCE_BENCH_DURATIOn"},
			},
		},
		Action: func(ctx *cli.Context) error {
			w := ansi.New()
			token, instance := auth.Account()
			var list v1.Site_List_Response
			err := klient.GET(
				context.Background(),
				instance+"/sites",
				&v1.Site_List_Request{},
				&list,
				token,
			)
			if err != nil {
				w.Err(err.Error).Exit()
			}
			m := make(map[string]struct{})
			for _, s := range list.List {
				m[s.Domain] = struct{}{}
			}
			args := ctx.Args()
			var g errgroup.Group
			gCtx, cancel := context.WithTimeout(context.Background(), ctx.Duration("duration"))
			defer cancel()
			users := ctx.Int("users")
			event := ctx.String("event")
			paths := ctx.StringSlice("paths")
			var stats []*Stats

			for i := 0; i < ctx.NArg(); i++ {
				a := args.Get(i)
				if _, ok := m[a]; !ok {
					w.Err(fmt.Sprintf("%q does not exist", a)).Flush()
					continue
				}
				stat := &Stats{
					Site:  a,
					Start: time.Now(),
				}
				stats = append(stats, stat)
				g.Go(bench(gCtx, stat, B{
					users:    users,
					instance: instance,
					site:     a,
					event:    event,
					paths:    paths,
				}))
			}
			g.Wait()
			summary(w, stats)
			return nil
		},
	}
}

type Stats struct {
	Site     string
	Users    atomic.Int64
	Requests atomic.Int64
	Start    time.Time
	End      time.Time
}

func summary(w *ansi.W, stats []*Stats) {
	for _, stat := range stats {
		w.Ok(stat.Site)
		users := stat.Users.Load()
		requests := stat.Requests.Load()
		duration := stat.End.Sub(stat.Start)
		rate := float64(requests) / duration.Seconds()
		w.KV("users", fmt.Sprint(users))
		w.KV("throughput", "%.2f req/s", rate).Flush()
	}
}

type B struct {
	users       int
	instance    string
	site, event string
	paths       []string
}

func bench(
	ctx context.Context,
	stats *Stats,
	b B,
) func() error {
	return func() error {
		defer func() {
			stats.End = time.Now()
		}()
		ips := ip(b.users)
		var g errgroup.Group
		for _, u := range ips {
			stats.Users.Add(1)
			g.Go(session(ctx, stats, S{
				instance: b.instance,
				usr:      u,
				domain:   b.site,
				agent:    agent(),
				event:    b.event,
				referrer: ref(),
				paths:    b.paths,
			}))
		}
		return g.Wait()
	}
}

var client = &http.Client{}

type S struct {
	instance, usr, domain, agent, event, referrer string
	paths                                         []string
}

func session(ctx context.Context,
	stats *Stats,
	s S,
) func() error {
	return func() error {
		e := entry.NewRequest()
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		for {
			select {
			case <-ctx.Done():
				return nil
			default:
				for _, path := range s.paths {
					stats.Requests.Add(1)
					e.EventName = s.event
					e.Domain = s.domain
					e.Referrer = s.referrer
					e.URI = "http://" + s.domain + path
					buf.Reset()
					enc.Encode(e)
					r, _ := http.NewRequest(http.MethodPost, s.instance+"/api/event", &buf)
					r.Header.Set("x-forwarded-for", s.usr)
					r.Header.Set("user-agent", s.agent)
					r.Header.Set("Accept", "application/json")
					r.Header.Set("content-type", "application/json")
					res, err := client.Do(r)
					if err != nil {
						slog.Error("failed sending client request", "domain", s.domain, "err", err.Error())
						continue
					}
					res.Body.Close()
					if res.StatusCode != http.StatusOK {
						slog.Error("unexpected response status code", "domain", s.domain, "code", res.StatusCode)
					}
				}
			}
		}
	}
}

func ip(n int) []string {
	n = min(n, len(geoip.IP))
	o := slices.Clone(geoip.IP)
	rand.Shuffle(len(o), func(i, j int) {
		o[i], o[j] = o[j], o[i]
	})
	return o[:n]
}

func ref() (o string) {
	n := rand.Intn(10)
	for k := range referrer.RefList {
		o = k
		if n < 1 {
			break
		}
		n--
	}
	p := strings.Split(o, ".")
	sort.Sort(sort.Reverse(referrer.StringSlice(p)))
	o = strings.Join(p, ".")
	return
}

func agent() string {
	return ua.AGENTS[rand.Intn(len(ua.AGENTS)-1)]
}
