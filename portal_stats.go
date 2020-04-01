package libwimark

import "time"

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
)

type dayTime struct {
	Year  int `json:"year" bson:"month"`
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

	UA     UserAgent `json:"ua" bson:"ua"`
	Vendor string    `json:"vendor" bson:"vendor"`
	Locale string    `json:"locale" bson:"locale"`

	Create   time.Time `json:"create" bson:"create"`
	CreateAt int64     `json:"create_at" bson:"create_at"`
}

type dailyProfileStat struct {
	Time     dayTime        `json:"time" bson:"time"`
	Profile  string         `json:"profile" bson:"profile"`
	Create   time.Time      `json:"create" bson:"create"`
	CreateAt int64          `json:"create_at" bson:"create_at"`
	Data     map[string]int `json:"data" bson:"data"`
}

type PortalDailyVendor = dailyProfileStat
type PortalDailyOS = dailyProfileStat
type PortalDailyLocale = dailyProfileStat
type PortalDailyType = dailyProfileStat
