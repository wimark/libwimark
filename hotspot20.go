package libwimark

type HS20_VenueName struct {
	Language string `json:"lang" bson:"lang"`
	Name     string `json:"name" bson:"name"`
}
type HS20_Venue struct {
	VenueGroup int              `json:"group" bson:"group"`
	VenueType  int              `json:"type" bson:"type"`
	Names      []HS20_VenueName `json:"names" bson:"names"`
}
type HS20_NAIAuth struct {
	Type    int `json:"type" bson:"type"`
	Subtype int `json:"subtype" bson:"subtype"`
}
type HS20_NAIRealm struct {
	Name string         `json:"name" bson:"name"`
	EAP  int            `json:"eap" bson:"eap"`
	Auth []HS20_NAIAuth `json:"auth" bson:"auth"`
}
type HS20_3GPPNetwork struct {
	Name string `json:"name" bson:"name"`
	MCC  int    `json:"mcc" bson:"mcc"`
	MNC  int    `json:"mnc" bson:"mnc"`
}
type HS20_Consortium struct {
	Name string `json:"name" bson:"name"`
	ID   string `json:"id" bson:"id"`
}
type Hotspot20Profile struct {
	Name        string             `json:"name" bson:"name"`
	NetType     int                `json:"net_type" bson:"net_type"`
	Internet    bool               `json:"internet" bson:"internet"`
	ASRA        bool               `json:"asra" bson:"asra"`
	ESR         bool               `json:"esr" bson:"esr"`
	UESA        bool               `json:"uesa" bson:"uesa"`
	HESSID      string             `json:"hessid" bson:"hessid"`
	Venue       HS20_Venue         `json:"venue" bson:"venue"`
	Domains     []string           `json:"domains" bson:"domains"`
	Realms      []HS20_NAIRealm    `json:"realms" bson:"realms"`
	Celluar     []HS20_3GPPNetwork `json:"celluar" bson:"celluar"`
	Consortiums []HS20_Consortium  `json:"consortiums" bson:"consortiums"`
}
