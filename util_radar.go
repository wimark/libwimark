package libwimark

import (
	"crypto/md5"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	TagMatch = "$match"
	TagGroup = "$group"
	TagLimit = "$limit"

	// resample periods
	RadarResampleHour  = "h"
	RadarResampleDay   = "d"
	RadarResampleWeek  = "w"
	RadarResampleMonth = "m"
)

// base object for
type RadarClientBaseObject struct {
	Id   string    `json:"id,omitempty" bson:"_id"`
	Date time.Time `json:"date" bson:"date"`
	Ts   int64     `json:"ts" bson:"ts"`
	MAC  string    `json:"mac" bson:"mac"`
	CPE  string    `json:"cpe" bson:"cpe"`
	Freq int       `json:"freq" bson:"freq"`
	RSSI int       `json:"rssi" bson:"rssi"`
	New  bool      `json:"new" bson:"new"`
}

type RadarClientIncome struct {
	Id   string    `json:"id,omitempty" bson:"_id"`
	Date time.Time `json:"date" bson:"date"`
	Ts   int64     `json:"ts" bson:"ts"`
	MAC  string    `json:"mac" bson:"mac"`
	CPE  string    `json:"cpe" bson:"cpe"`
}

type RadarClientVisit struct {
	Id        string    `json:"id,omitempty" bson:"_id"`
	MAC       string    `json:"mac" bson:"mac"`
	DateStart time.Time `json:"start" bson:"start"`
	DateStop  time.Time `json:"stop" bson:"stop"`
	Duration  int64     `json:"duration" bson:"duration"`
	CPE       string    `json:"cpe,omitempty" bson:"cpe"`
	New       bool      `json:"new" bson:"new"`
	RSSI      int       `json:"rssi" bson:"rssi"`
}

type RadarClientVisitor struct {
	Id          string             `json:"id,omitempty" bson:"_id"`
	MAC         string             `json:"mac" bson:"mac"`
	DateFirst   time.Time          `json:"visit_first" bson:"visit_first"`
	Duration    int64              `json:"duration" bson:"duration"`
	DurationAvg int64              `json:"duration_avg" bson:"duration_avg"`
	VisitsCount int                `json:"visits_count" bson:"visits_count"`
	Visits      []RadarClientVisit `json:"visits" bson:"visits"`
	New         bool               `json:"new" bson:"new"`
}

type RadarClientFirst struct {
	Id   string    `json:"id,omitempty" bson:"_id"`
	Date time.Time `json:"date" bson:"date"`
	Ts   int64     `json:"ts" bson:"ts"`
	MAC  string    `json:"mac" bson:"mac"`
	CPE  string    `json:"cpe" bson:"cpe"`
}

type RadarClientVisitGroup struct {
	All         int64 `json:"all" bson:"all"`
	New         int64 `json:"new" bson:"new"`
	DurationAvg int64 `json:"duration_avg" bson:"duration_avg"`
}

type AnalyticsMwHttpRequest struct {
	CPEs         []string `query:"cpes[]"`
	Location     string   `query:"location"`
	WithChildren bool     `query:"with_childs"`

	Start   int64  `query:"start"`
	Stop    int64  `query:"stop"`
	Timeout int    `query:"timeout"`
	Period  string `query:"period"`

	Rate             int             `query:"rate"`
	Raw              bool            `query:"raw"`
	Long             bool            `query:"long"`
	Hash             bool            `query:"hash"`
	Filter           string          `query:"filter"`
	Duration         int             `query:"duration"`
	WithNightClients bool            `query:"with_night_clients"`
	WorkTimeRange    string          `query:"work_time_range"`
	Zone             string          `query:"zone"`
	ExportMac        RadarExportMacs `query:"export_mac"`
}

func (r *AnalyticsMwHttpRequest) String() string {
	s := fmt.Sprintf(
		"start=%d&stop=%d&timeout=%d&period=%s&rate=%d&filter=%s&duration=%d&raw=%s&long=%s&hash=%s&with_night_clients=%s&work_time_range=%s&zone=%s&export_mac=%s",
		r.Start, r.Stop, r.Timeout, r.Period, r.Rate, r.Filter, r.Duration,
		strconv.FormatBool(r.Raw),
		strconv.FormatBool(r.Long),
		strconv.FormatBool(r.Hash),
		strconv.FormatBool(r.WithNightClients),
		r.WorkTimeRange, r.Zone, r.ExportMac.String())
	if len(r.CPEs) != 0 {
		cpes := strings.Join(r.CPEs, "&cpes[]=")
		s = s + "&cpes[]=" + cpes
	}
	return s
}

type AnalyticsMwHttpResponse struct {
	Status string                 `json:"status"`
	Desc   string                 `json:"error,omitempty"`
	Data   map[string]interface{} `json:"data"`
}

type RadarExportPeriod struct {
	Start int64 `json:"start" bson:"start"`
	Stop  int64 `json:"stop" bson:"stop"`
}

type RadarExportCreds struct {
	// subject for email subject
	Subject string `json:"subject" bson:"subject"`

	// share for list of emails, yandex accounts to share with or CTN for BeePro integration
	Share []string `json:"share" bson:"share"`
}

type RadarExportState struct {
	Id                     int    `json:"id" bson:"id"`
	Type                   string `json:"type" bson:"type"`
	Status                 string `json:"status" bson:"status"`
	ItemQuantity           int    `json:"item_quantity" bson:"item_quantity"`
	ValidUniqueQuantity    int    `json:"valid_unique_quantity"  bson:"valid_unique_quantity"`
	MatchedQuantity        int    `json:"matched_quantity"  bson:"matched_quantity"`
	CookiesMatchedQuantity int    `json:"cookies_matched_quantity"  bson:"cookies_matched_quantity"`
}

type RadarExportObject struct {
	Id string `json:"id" bson:"_id"`

	// common settings
	Name string `json:"name" bson:"name"`
	Desc string `json:"desc" bson:"desc"`

	// config part
	CPEs       []string          `json:"cpes" bson:"cpes"`
	Type       RadarExportType   `json:"type" bson:"type"`
	Creds      RadarExportCreds  `json:"creds" bson:"creds"`
	Format     RadarExportFormat `json:"format" bson:"format"`
	Period     RadarExportPeriod `json:"period" bson:"period"`
	Filter     RadarExportFilter `json:"filter" bson:"filter"`
	MacFilter  RadarExportMacs   `json:"mac_filter" bson:"mac_filter"`
	Hash       bool              `json:"hash" bson:"hash"`
	Auto       bool              `json:"auto" bson:"auto"`               // auto export
	AutoPeriod int               `json:"auto_period" bson:"auto_period"` // auto period in hours. 24 for 1-day update

	// state part
	CreateAt int64             `json:"create_at" bson:"create_at"`
	LastAt   int64             `json:"last_at" bson:"last_at"`
	Status   RadarExportStatus `json:"status" bson:"status"`
	State    RadarExportState  `json:"state" bson:"state"`

	// CTN for BeePro -- just for storing in DB
	CTN string `json:"ctn" bson:"ctn"`
}

type RadarExportObjectUpdate struct {
	Name string `json:"name" bson:"name"`
	Desc string `json:"desc" bson:"desc"`
}

type RadarExportResult struct {
	Id   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`

	CreateAt time.Time `json:"create_at" bson:"create_at"`
	Append   bool      `json:"append" bson:"append"`
	CPEs     []string  `json:"cpes" bson:"cpes"`
	MACs     []string  `json:"macs" bson:"macs"`
}

type RadarExportUpdate struct {
	Type string   `json:"type" bson:"type"`
	IDs  []string `json:"ids" bson:"ids"`
}

func MacAddrShrink(s string) string {
	return strings.ToLower(stripchars(s, ":-."))
}

func MacAddrIsGlobalAssigned(s string) bool {
	if len(s) < 2 {
		return false
	}

	first_byte, _ := strconv.ParseInt(s[0:2], 16, 64)

	return first_byte&2 == 0
}

func MacAddrVendor(s string) string {
	if len(s) < 6 {
		return ""
	}
	return MACPrefixVendorMap.Get(s[0:6])
}

func MacAddrIsReal(s string) bool {
	return MacAddrIsGlobalAssigned(s) && MacAddrVendor(s) != ""
}

func MacAddrHash(mac string) string {
	if len(mac) < 12 {
		return ""
	}
	m, _ := net.ParseMAC(fmt.Sprintf("%s.%s.%s", mac[0:4],
		mac[4:8], mac[8:12]))
	return fmt.Sprintf("%X", md5.Sum(m))
}

func stripchars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if !strings.ContainsRune(chr, r) {
			return r
		}
		return -1
	}, str)
}
