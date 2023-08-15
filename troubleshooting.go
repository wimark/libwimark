package libwimark

type Troubleshooting struct {
	State                  RadioActiveState       `json:"state" bson:"state"`
	TroubleshootingFilters TroubleshootingFilters `json:"troubleshooting_filters" bson:"troubleshooting_filters"`
}

type TroubleshootingFilters []TroubleshootingFilter

type TroubleshootingFilter struct {
	ID       string `json:"id" bson:"_id"`
	MacIP    string `json:"mac_ip" bson:"mac_ip"`
	FileName string `json:"filename" bson:"filename"`
	URL      string `json:"url" bson:"url"`
	StartAt  int64  `json:"start_at" bson:"start_at"`
	StopAt   string `json:"stop_at" bson:"stop_at"`
}