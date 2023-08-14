package libwimark

import "time"

type NTPGeneral struct {
	Active        NTPGeneralActive `json:"active" bson:"active"`
	LocalTimeTZ   time.Time        `json:"local_time_tz" bson:"local_time_tz"`
	LocalTimeZone string           `json:"local_time_zone" bson:"local_time_zone"`
	ExternalNTP   string           `json:"external_ntp" bson:"external_ntp"`
	NTPPort       string           `json:"ntp_port" bson:"ntp_port"`
	NTPServers    NTPServers       `json:"ntp_servers" bson:"ntp_servers"`
}

type NTPServers []NTPServer

type NTPServer struct {
	ID             string          `json:"id" bson:"_id"`
	Title          string          `json:"title" bson:"title"`
	Address        string          `json:"address" bson:"address"`
	ResolveAddress string          `json:"-" bson:"resolve_address"`
	Port           string          `json:"port" bson:"port"`
	Status         NTPServerStatus `json:"status" bson:"status"`
	Priority       string          `json:"priority" bson:"priority"`
}
