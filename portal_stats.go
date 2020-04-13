package libwimark

import (
	"fmt"
	"time"
)

const (
	// CollPortalClientLog collection with portal client state moves
	CollPortalClientLog = "portal_client_log"
	// CollPortalClientStat collection with clean portal pass data
	CollPortalClientStat = "portal_client_stat"
	// CollPortalDailyVendor collection with daily aggregated vendors
	CollPortalDailyVendor = "portal_daily_vendor"
	// CollPortalDailyOS collection with daily aggregated OS
	CollPortalDailyOS = "portal_daily_os"
	// CollPortalDailyLocale daily aggregated with locales
	CollPortalDailyLocale = "portal_daily_locale"
	// CollPortalDailyType daily aggregated with useragent type
	CollPortalDailyType = "portal_daily_type"

	CollPortalDailyFirstVisit = "portal_daily_first_visit"
	CollPortalDailyAuth       = "portal_daily_auth"
	CollPortalDailyAuthen     = "portal_daily_authen"
)

// DayTime struct fot represent year-month-day object
type DayTime struct {
	Year  int `json:"year" bson:"year"`
	Month int `json:"month" bson:"month"`
	Day   int `json:"day" bson:"day"`
}

// PortalClientLog for logging every user state (index should be for 1 month)
type PortalClientLog struct {
	Profile   string    `json:"profile" bson:"profile"`
	MAC       string    `json:"mac" bson:"mac"`
	Identity  string    `json:"identity,omitempty" bson:"identity,omitempty"`
	Account   string    `json:"account,omitempty" bson:"account"`
	UserAgent string    `json:"useragent" bson:"useragent"`
	Locale    string    `json:"locale" bson:"locale"`
	Path      string    `json:"path" bson:"path"`
	State     string    `json:"state" bson:"state"`
	Success   bool      `json:"success" bson:"success"`
	Error     string    `json:"error,omitempty" bson:"error,omitempty"`
	Create    time.Time `json:"create" bson:"create"`
	CreateAt  int64     `json:"create_at" bson:"create_at"`
}

// PortalClientStat for stats portal passed users
type PortalClientStat struct {
	Profile  string `json:"profile" bson:"profile"`
	MAC      string `json:"mac" bson:"mac"`
	Identity string `json:"identity,omitempty" bson:"identity,omitempty"`
	Account  string `json:"account,omitempty" bson:"account"`

	AuthenType string `json:"authen" bson:"authen,omitempty"`
	AuthType   string `json:"auth" bson:"auth,omitempty"`

	UA         UserAgent `json:"ua" bson:"ua"`
	Vendor     string    `json:"vendor" bson:"vendor"`
	Locale     string    `json:"locale" bson:"locale"`
	FirstVisit string    `json:"first_visit" bson:"first_visit"`

	Create   time.Time `json:"create" bson:"create"`
	CreateAt int64     `json:"create_at" bson:"create_at"`
}

// DailyProfileStat struct represent daily aggregated stats
// common for os, vendor, devices, etc
type DailyProfileStat struct {
	Create   time.Time `json:"create" bson:"create"`
	CreateAt int64     `json:"create_at" bson:"create_at"`
	Time     DayTime   `json:"time" bson:"time"`

	Profile string   `json:"profile" bson:"profile"`
	Values  []string `json:"values" bson:"values"`
	Counts  []int    `json:"counts" bson:"counts"`
}

// DailyProfileStatDB struct wrapper with ID for DailyProfileStat
type DailyProfileStatDB struct {
	ID     string           `json:"id" bson:"_id"`
	Object DailyProfileStat `json:",inline" bson:",inline"`
}

// NewDailyProfileStat func return new object DailyProfileStat
func NewDailyProfileStatDBYesterday(t time.Time, profile string,
	values []string, counts []int) DailyProfileStatDB {
	var tYesterday = t.AddDate(0, 0, -1)
	var year, month, day = tYesterday.Year(), int(tYesterday.Month()), tYesterday.Day()
	return DailyProfileStatDB{
		ID: fmt.Sprintf("%d-%d-%d-%s", year, month, day, profile),
		Object: DailyProfileStat{
			Create:   t,
			CreateAt: t.Unix(),
			Time: DayTime{
				Year:  year,
				Month: month,
				Day:   day,
			},
			Profile: profile,
			Values:  values,
			Counts:  counts,
		},
	}
}

type PortalDailyVendor = DailyProfileStat
type PortalDailyOS = DailyProfileStat
type PortalDailyLocale = DailyProfileStat
type PortalDailyType = DailyProfileStat
