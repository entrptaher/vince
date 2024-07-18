package api

import (
	"context"
	"net/http"
	"strings"
	"time"

	v1 "github.com/vinceanalytics/vince/gen/go/vince/v1"
	"github.com/vinceanalytics/vince/internal/logger"
	"github.com/vinceanalytics/vince/internal/timeutil"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Params represents the query parameters passed in an HTTP request.
// Query parameter keys are case-sensitive, as per the HTTP spec.
type Params map[string]string

// NewParams returns a new Params from the given HTTP request.
func NewParams(r *http.Request) Params {
	qp := make(Params)

	for k, v := range r.URL.Query() {
		qp[k] = v[0]
	}

	if username, password, ok := r.BasicAuth(); ok {
		qp["username"] = username
		qp["password"] = password
	}
	if token, ok := Bearer(r.Header.Get("Authorization")); ok {
		qp["token"] = token
	}
	return qp
}

func Bearer(header string) (string, bool) {
	if header == "" {
		return "", false
	}
	return strings.CutPrefix(header, "Bearer ")
}

func (qp Params) IsBearer() bool {
	_, ok := qp["token"]
	return ok
}

func (qp Params) BearerToken() string {
	return qp["token"]
}

func (qp Params) Basic() (username, password string, ok bool) {
	username, ok = qp["username"]
	password = qp["password"]
	return
}

func (qp Params) SiteID() string {
	return qp["site_id"]
}

// Avoid leaking internal errors to client. The actual error is logged and this
// is returned back to the client.
var InternalError = status.Error(codes.Internal, "Something went wrong")

func (p Params) PeriodToRange(ctx context.Context, now func() time.Time, period *v1.TimePeriod) (start, end time.Time) {
	date := p.parseDate(ctx, now)
	switch e := period.Value.(type) {
	case *v1.TimePeriod_Base_:
		switch e.Base {
		case v1.TimePeriod_day:
			end = date
			start = timeutil.BeginDay(end)
		case v1.TimePeriod__7d:
			end = date
			start = end.AddDate(0, 0, -6)
		case v1.TimePeriod__30d:
			end = date
			start = end.AddDate(0, 0, -30)
		case v1.TimePeriod_mo:
			end = date
			start = timeutil.BeginMonth(end)
			end = timeutil.EndMonth(end)
		case v1.TimePeriod__6mo:
			end = timeutil.EndMonth(date)
			start = timeutil.BeginMonth(end.AddDate(0, -5, 0))
		case v1.TimePeriod__12mo:
			end = timeutil.EndMonth(date)
			start = timeutil.BeginMonth(end.AddDate(0, -11, 0))
		case v1.TimePeriod_year:
			end = timeutil.EndYear(date)
			start = timeutil.BeginYear(end)
		}

	case *v1.TimePeriod_Custom_:
		end = e.Custom.End.AsTime()
		start = e.Custom.Start.AsTime()
	}
	return
}

func ValidByPeriod(period *v1.TimePeriod, i v1.Interval) bool {
	switch e := period.Value.(type) {
	case *v1.TimePeriod_Base_:
		switch e.Base {
		case v1.TimePeriod_day:
			return i == v1.Interval_minute || i == v1.Interval_hour
		case v1.TimePeriod__7d:
			return i == v1.Interval_hour || i == v1.Interval_date
		case v1.TimePeriod_mo, v1.TimePeriod__30d:
			return i == v1.Interval_date || i == v1.Interval_week
		case v1.TimePeriod__6mo, v1.TimePeriod__12mo, v1.TimePeriod_year:
			return i == v1.Interval_date || i == v1.Interval_week || i == v1.Interval_month
		default:
			return false
		}
	case *v1.TimePeriod_Custom_:
		return i == v1.Interval_date || i == v1.Interval_week || i == v1.Interval_month
	default:
		return false
	}
}

func (p Params) Period(ctx context.Context) *v1.TimePeriod {
	value := p["period"]
	switch value {
	case "12mo":
		return &v1.TimePeriod{Value: &v1.TimePeriod_Base_{Base: v1.TimePeriod__12mo}}
	case "6mo":
		return &v1.TimePeriod{Value: &v1.TimePeriod_Base_{Base: v1.TimePeriod__6mo}}
	case "month":
		return &v1.TimePeriod{Value: &v1.TimePeriod_Base_{Base: v1.TimePeriod_mo}}
	case "30day":
		return &v1.TimePeriod{Value: &v1.TimePeriod_Base_{Base: v1.TimePeriod__30d}}
	case "7day":
		return &v1.TimePeriod{Value: &v1.TimePeriod_Base_{Base: v1.TimePeriod__7d}}
	case "day":
		return &v1.TimePeriod{Value: &v1.TimePeriod_Base_{Base: v1.TimePeriod_day}}
	case "custom":
		date := p["date"]
		if date == "" {
			logger.Get(ctx).Error("custom period specified with missing date")
			return nil
		}
		from, to, _ := strings.Cut(date, ",")

		start, err := time.Parse(time.DateOnly, from)
		if err != nil {
			logger.Get(ctx).Error("Invalid date for custom period", "date", date, "err", err)
			return nil
		}
		end, err := time.Parse(time.DateOnly, to)
		if err != nil {
			logger.Get(ctx).Error("Invalid date for custom period", "date", date, "err", err)
			return nil
		}
		return &v1.TimePeriod{
			Value: &v1.TimePeriod_Custom_{
				Custom: &v1.TimePeriod_Custom{
					Start: timestamppb.New(start),
					End:   timestamppb.New(end),
				},
			},
		}
	default:
		return &v1.TimePeriod{
			Value: &v1.TimePeriod_Base_{Base: v1.TimePeriod__30d},
		}
	}
}

func (p Params) parseDate(ctx context.Context, now func() time.Time) time.Time {
	date := p["date"]
	if date == "" {
		return timeutil.EndDay(now())
	}
	v, err := time.Parse(time.DateOnly, date)
	if err != nil {
		fall := timeutil.EndDay(now())
		logger.Get(ctx).Error("failed parsing date falling back to now",
			"date", date, "now", fall.Format(time.DateOnly), "err", err)
		return fall
	}
	return v
}

func (p Params) Metrics(ctx context.Context) (o []v1.Metric) {
	metrics := p["metrics"]
	if metrics == "" {
		return []v1.Metric{v1.Metric_visitors}
	}
	for _, m := range strings.Split(metrics, ",") {
		v, ok := v1.Metric_value[m]
		if !ok {
			logger.Get(ctx).Error("Skipping unexpected metric name", "metric", m)
			continue
		}
		o = append(o, v1.Metric(v))
	}
	return
}

func (p Params) Property(ctx context.Context) (o []v1.Property) {
	props := p["property"]
	for _, m := range strings.Split(props, ",") {
		v, ok := v1.Property_value[m]
		if !ok {
			logger.Get(ctx).Error("Skipping unexpected property name", "property", m)
			continue
		}
		o = append(o, v1.Property(v))
	}
	return
}

func (p Params) Interval(ctx context.Context) v1.Interval {
	i := p["interval"]
	if i == "" {
		return v1.Interval_date
	}
	v, ok := v1.Interval_value[i]
	if !ok {
		if i != "" {
			logger.Get(ctx).Error("Skipping unexpected interval value", "interval", i)
		}
	}
	return v1.Interval(v)
}

func (p Params) Filters(ctx context.Context) (o []*v1.Filter) {
	fs := p["filters"]
	if fs == "" {
		return
	}
	for _, f := range strings.Split(fs, ",") {
		key, value, op, ok := sep(f)
		if !ok {
			logger.Get(ctx).Error("Skipping unexpected filter ", "filter", f)
			continue
		}
		p, ok := v1.Property_value[key]
		if !ok {
			logger.Get(ctx).Error("Skipping unexpected filter property ", "filter", f, "property", key)
			continue
		}
		o = append(o, &v1.Filter{
			Property: v1.Property(p),
			Op:       op,
			Value:    value,
		})
	}
	return
}

func sep(f string) (key, value string, op v1.Filter_OP, ok bool) {
	key, value, ok = strings.Cut(f, "==")
	if ok {
		op = v1.Filter_equal
		return
	}
	key, value, ok = strings.Cut(f, "!=")
	if ok {
		op = v1.Filter_not_equal
		return
	}
	key, value, ok = strings.Cut(f, "~=")
	if ok {
		op = v1.Filter_re_equal
		return
	}
	key, value, ok = strings.Cut(f, "!~")
	if ok {
		op = v1.Filter_re_not_equal
		return
	}
	return
}
