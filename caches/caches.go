package caches

import (
	"context"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/gernest/vince/models"
	"golang.org/x/time/rate"
)

type sessionKey struct{}
type sitesKey struct{}

func Open(ctx context.Context) (context.Context, error) {
	session, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     2 << 20,
		BufferItems: 64,
	})
	if err != nil {
		return nil, err
	}
	sites, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     1 << 20,
		BufferItems: 64,
	})
	if err != nil {
		session.Close()
		return nil, err
	}
	ctx = context.WithValue(ctx, sessionKey{}, session)
	ctx = context.WithValue(ctx, sitesKey{}, sites)
	return ctx, nil
}

func Close(ctx context.Context) {
	Session(ctx).Close()
	Site(ctx).Close()
}

func Session(ctx context.Context) *ristretto.Cache {
	return ctx.Value(sessionKey{}).(*ristretto.Cache)
}

func Site(ctx context.Context) *ristretto.Cache {
	return ctx.Value(sitesKey{}).(*ristretto.Cache)
}

type SiteRate struct {
	SID  uint64
	UID  uint64
	Rate *rate.Limiter
}

func SetSite(ctx context.Context, ttl time.Duration) func(*models.CachedSite) {
	cache := Site(ctx)
	return func(cs *models.CachedSite) {
		cache.SetWithTTL(cs.Domain, &SiteRate{
			SID:  cs.ID,
			UID:  cs.UserID,
			Rate: rate.NewLimiter(cs.RateLimit()),
		}, 1, ttl)
	}
}

func AllowSite(ctx context.Context, domain string) (uid, sid uint64, ok bool) {
	x, _ := Site(ctx).Get(domain)
	if x != nil {
		r := x.(*SiteRate)
		return r.UID, r.SID, r.Rate.Allow()
	}
	return
}
