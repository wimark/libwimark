package libwimark

const (
	CollPortalClientLog = "portal_client_log"
)

type dayTime struct {
	Year  int `json:"year" bson:"month"`
	Month int `json:"month" bson:"month"`
	Day   int `json:"day" bson:"day"`
}

// PortalClientLog for logging every user state
type PortalClientLog struct {
	TS        int64  `json:"ts" bson:"ts"`
	Profile   string `json:"profile" bson:"profile"`
	MAC       string `json:"mac" bson:"mac"`
	Identity  string `json:"identity,omitempty" bson:"identity,omitempty"`
	Account   string `json:"account,omitempty" bson:"account"`
	UserAgent string `json:"useragent" bson:"useragent"`
	Locale    string `json:"locale" bson:"locale"`
	Path      string `json:"path" bson:"path"`
	State     string `json:"state" bson:"state"`
	Success   bool   `json:"success" bson:"success"`
	Error     string `json:"error,omitempty" bson:"error,omitempty"`
}
