package timeseries

type Property uint8

const (
	Base           Property = 0
	Event          Property = 1
	Page           Property = 2
	EntryPage      Property = 3
	ExitPage       Property = 4
	Referrer       Property = 5
	UtmMedium      Property = 6
	UtmSource      Property = 7
	UtmCampaign    Property = 8
	UtmContent     Property = 9
	UtmTerm        Property = 10
	UtmDevice      Property = 11
	UtmBrowser     Property = 12
	BrowserVersion Property = 13
	Os             Property = 14
	OsVersion      Property = 15
	Country        Property = 16
	Region         Property = 17
	City           Property = 18
)

// Enum value maps for Property.
var (
	_prop_name = map[uint8]string{
		0:  "base",
		1:  "event",
		2:  "page",
		3:  "entryPage",
		4:  "exitPage",
		5:  "referrer",
		6:  "UtmMedium",
		7:  "UtmSource",
		8:  "UtmCampaign",
		9:  "UtmContent",
		10: "UtmTerm",
		11: "UtmDevice",
		12: "UtmBrowser",
		13: "browserVersion",
		14: "os",
		15: "osVersion",
		16: "country",
		17: "region",
		18: "city",
	}
	_prop_value = map[string]uint8{
		"base":           0,
		"event":          1,
		"page":           2,
		"entryPage":      3,
		"exitPage":       4,
		"referrer":       5,
		"UtmMedium":      6,
		"UtmSource":      7,
		"UtmCampaign":    8,
		"UtmContent":     9,
		"UtmTerm":        10,
		"UtmDevice":      11,
		"UtmBrowser":     12,
		"browserVersion": 13,
		"os":             14,
		"osVersion":      15,
		"country":        16,
		"region":         17,
		"city":           18,
	}
)

func (p Property) String() string {
	return _prop_name[uint8(p)]
}

type Metric uint8

const (
	Visitors      Metric = 0
	Views         Metric = 1
	Events        Metric = 2
	Visits        Metric = 3
	BounceRate    Metric = 4
	VisitDuration Metric = 5
	ViewsPerVisit Metric = 6
)

// Enum value maps for Metric.
var (
	_metric_name = map[uint8]string{
		0: "visitors",
		1: "views",
		2: "events",
		3: "visits",
		4: "bounceRate",
		5: "visitDuration",
		6: "viewsPerVisit",
	}
	_metric_value = map[string]uint8{
		"visitors":      0,
		"views":         1,
		"events":        2,
		"visits":        3,
		"bounceRate":    4,
		"visitDuration": 5,
		"viewsPerVisit": 6,
	}
)

func (m Metric) String() string {
	return _metric_name[uint8(m)]
}
