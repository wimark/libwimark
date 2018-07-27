package libwimark

type I18nName struct {
	Language string `json:"lang" bson:"lang"`
	Name     string `json:"name" bson:"name"`
}
type HS20_Venue struct {
	VenueGroup int        `json:"group" bson:"group"`
	VenueType  int        `json:"type" bson:"type"`
	Names      []I18nName `json:"names" bson:"names"`
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
type HS20_IPTypes struct {
	IpV4Type int `json:"ipv4type" bson:"ipv4type"`
	IpV6Type int `json:"ipv6type" bson:"ipv6type"`
}
type HS20_ConnCaps struct {
	Protocol int `json:"protocol" bson:"protocol"`
	Port     int `json:"port" bson:"port"`
	Status   int `json:"status" bson:"status"`
}
type HS20_WanMetrics struct {
	WANInfo         int `json:"wan_info" bson:"wan_info"`
	DownlinkSpeed   int `json:"dl_speed" bson:"dl_speed"`
	UplinkSpeed     int `json:"ul_speed" bson:"ul_speed"`
	DownlinkLoad    int `json:"dl_load" bson:"dl_load"`
	UplinkLoad      int `json:"ul_load" bson:"ul_load"`
	MeasureDuration int `json:"lmd" bson:"lmd"`
}
type Hotspot20Profile struct {
	Name        string             `json:"name" bson:"name"`
	NetType     int                `json:"net_type" bson:"net_type"`
	Internet    bool               `json:"internet" bson:"internet"`
	ASRA        bool               `json:"asra" bson:"asra"`
	ESR         bool               `json:"esr" bson:"esr"`
	UESA        bool               `json:"uesa" bson:"uesa"`
	DGAF        bool               `json:"dgaf" bson:"dgaf"`
	HESSID      string             `json:"hessid" bson:"hessid"`
	Venue       HS20_Venue         `json:"venue" bson:"venue"`
	Domains     []string           `json:"domains" bson:"domains"`
	Realms      []HS20_NAIRealm    `json:"realms" bson:"realms"`
	Cellular    []HS20_3GPPNetwork `json:"cellular" bson:"cellular"`
	Consortiums []HS20_Consortium  `json:"consortiums" bson:"consortiums"`
	OperNames   []I18nName         `json:"oper_names" bson:"oper_names"`
	IpTypes     HS20_IPTypes       `json:"ip_types" bson:"ip_types"`
	ConnCaps    []HS20_ConnCaps    `json:"conn_caps" bson:"conn_caps"`
	WanMetrics  HS20_WanMetrics    `json:"wan_metrics" bson:"wan_metrics"`
	OperClasses []int              `json:"oper_classes" bson:"oper_classes"`
}
