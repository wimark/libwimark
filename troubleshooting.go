package libwimark

type RadioactiveLog struct {
	StartDate  int64    `json:"start_date" bson:"start_date"`
	StopDate   int64    `json:"stop_date" bson:"stop_date"`
	MACAddress []string `json:"mac_address" bson:"mac_address"`
	DebugDHCP  string   `json:"debug_dhcp" bson:"debug_dhcp"`
	DebugNTP   string   `json:"debug_ntp" bson:"debug_ntp"`
	Hostap     string   `json:"hostap" bson:"hostap"`
}
