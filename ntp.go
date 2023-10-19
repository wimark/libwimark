package libwimark

import (
	"time"
)

type NTPServers []NTPServer

type NTPServer struct {
	ID       string          `json:"id" bson:"_id"`
	Title    string          `json:"title" bson:"title"`
	Address  string          `json:"address" bson:"address"`
	Port     string          `json:"port" bson:"port"`
	Status   NTPServerStatus `json:"status" bson:"status"`
	Priority string          `json:"priority" bson:"priority"`
	Prefer   bool            `json:"prefer" bson:"prefer"`
	LastSync time.Time       `json:"last_sync,omitempty" bson:"last_sync"`
}

type NtpTimeZone struct {
	ID          string            `json:"id" bson:"_id"`
	Zone        string            `json:"time_zone" bson:"time_zone"`
	Offset      NtpTimeZoneOffset `json:"offset" bson:"offset"`
	LastUpdated time.Time         `json:"last_updated" bson:"last_updated"`
}

type NtpTimeZoneOffset struct {
	Hour int `json:"hour" bson:"hour"`
	Min  int `json:"min" bson:"min"`
}

type NtpServerDeps struct {
	Title    string
	Host     string
	Port     string
	Priority string
	Prefer   bool
}

func NewNtpServer(deps NtpServerDeps) *NTPServer {
	return &NTPServer{
		ID:       NewUUID(),
		Title:    deps.Title,
		Address:  deps.Host,
		Port:     deps.Port,
		Priority: deps.Priority,
		Prefer:   deps.Prefer,
	}
}

func NewNtpTimeZone(zone string, hour, min int) *NtpTimeZone {
	return &NtpTimeZone{
		ID:          NewUUID(),
		Zone:        zone,
		Offset:      NtpTimeZoneOffset{Hour: hour, Min: min},
		LastUpdated: time.Now(),
	}
}
