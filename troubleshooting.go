package libwimark

// JSONRPC functions for configurer
const (
	JSONRPC_TROUBLESHOOTING = "troubleshooting"
)

// for cpeagent:*
type TroubleshootingCpe struct {
	State RadioActiveState `json:"state"`
}

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
	StopAt   int64  `json:"stop_at" bson:"stop_at"`
}

type TroubleshootingGenerateLog struct {
	ID       string `json:"id"`
	Interval int64  `json:"interval"`
}
