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
	TAG_MATCH = "$match"
	TAG_GROUP = "$group"
	TAG_LIMIT = "$limit"

	// необработанные пробы от точки -- TTL 24 часа
	COLL_RADAR_BASE = "radar_probes_raw"
	// отфильтрованные пробы по реальным MAC -- TTL 24 часа
	COLL_RADAR_BASE_REAL = "radar_probes_real"
	// первое появление конкретного клиента с привязской к точке -- TTL 1 год
	COLL_RADAR_VISITS_FIRST = "radar_visits_first"

	// подсчет клиентских визитов RadarClientVisit -- TTL 1 год -- аггрегация раз в сутки
	COLL_RADAR_VISITS = "radar_visits"
	// клиентские визиты с привязкой к конкретному часу -- TTL 1 год -- аггрегация раз в сутки
	COLL_RADAR_VISITS_HOUR = "radar_visits_hour"

	RADAR_RESAMPLE_HOUR  = "h"
	RADAR_RESAMPLE_DAY   = "d"
	RADAR_RESAMPLE_WEEK  = "w"
	RADAR_RESAMPLE_MONTH = "m"
)

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

type AnalyticsMwHttpRequest struct {
	CPEs     []string `query:"cpes"`
	Location string   `query:"location"`

	Start   int64  `query:"start"`
	Stop    int64  `query:"stop"`
	Timeout int    `query:"timeout"`
	Period  string `query:"period"`

	Rate int  `query:"rate"`
	Raw  bool `query:"raw"`
	Long bool `query:"long"`
	Hash bool `query:"hash"`
}

func (r *AnalyticsMwHttpRequest) String() string {
	s := fmt.Sprintf("start=%d&stop=%d&timeout=%d&period=%s&rate=%d&raw=%s&long=%s&hash=%s",
		r.Start, r.Stop, r.Timeout, r.Period, r.Rate,
		strconv.FormatBool(r.Raw),
		strconv.FormatBool(r.Long),
		strconv.FormatBool(r.Hash))
	if len(r.CPEs) != 0 {
		cpes := strings.Join(r.CPEs, ",")
		s = s + "&cpes=" + cpes
	}
	return s
}

type AnalyticsMwHttpResponse struct {
	Status string                 `json:"status"`
	Desc   string                 `json:"error,omitempty"`
	Data   map[string]interface{} `json:"data"`
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
	v, _ := ManufacturerMap[s[0:6]]
	return v
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
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}
