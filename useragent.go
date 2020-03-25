package libwimark

// UserAgent struct for parse useragent to internal struct
type UserAgent struct {
	Type      string `json:"type" bson:"type"`
	Device    string `json:"device" bson:"device"`
	Name      string `json:"name" bson:"name"`
	OS        string `json:"os" bson:"os"`
	OSVersion string `json:"os_version" bson:"os_version"`
}
