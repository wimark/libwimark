package libwimark

import (
	"time"
)

// Deprecated: use only backend
type NTPGeneral struct {
	Active        NTPGeneralActive `json:"active" bson:"active"`
	LocalTimeTZ   time.Time        `json:"local_time_tz" bson:"local_time_tz"`
	LocalTimeZone string           `json:"local_time_zone" bson:"local_time_zone"`
	ExternalNTP   string           `json:"external_ntp" bson:"external_ntp"`
	NTPPort       string           `json:"ntp_port" bson:"ntp_port"`
	NTPServers    NTPServers       `json:"ntp_servers" bson:"ntp_servers"`
}

// Deprecated: use only backend
type NTPServers []NTPServer

// Deprecated: use only backend
type NTPServer struct {
	ID             string          `json:"id" bson:"_id"`
	Title          string          `json:"title" bson:"title"`
	Address        string          `json:"address" bson:"address"`
	ResolveAddress string          `json:"resolve_address" bson:"resolve_address"`
	Port           string          `json:"port" bson:"port"`
	Status         NTPServerStatus `json:"status" bson:"status"`
	Priority       string          `json:"priority" bson:"priority"`
}

type NtpServer struct {
	ID      string           `json:"id" bson:"_id"`
	Title   string           `json:"title" bson:"title"`
	Network NtpServerNetwork `json:"network" bson:"network"`
	Meta    NtpServerMeta    `json:"meta" bson:"meta"`
}

type NtpServerMeta struct {
	Priority uint32          `json:"priority" bson:"priority"`
	Status   NtpServerStatus `json:"status" bson:"status"`
}

type NtpServerNetwork struct {
	IP   string `json:"ip" bson:"ip"`
	Host string `json:"host" bson:"host"`
	Port string `json:"port" bson:"port"`
}

func NewNtpTimeSettings(sec int64, zone string) *NtpTimeSettings {
	return &NtpTimeSettings{
		ID:          NewUUID(),
		Time:        sec,
		TimeZone:    zone,
		LastUpdated: time.Now(),
	}
}

type NtpTimeSettings struct {
	ID          string    `json:"id" bson:"_id"`
	Time        int64     `json:"time" bson:"time"`
	TimeZone    string    `json:"time_zone" bson:"time_zone"`
	LastUpdated time.Time `json:"last_updated" bson:"last_updated"`
}
