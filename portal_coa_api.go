package libwimark

// Platform part
// --------

// CoARequest - request payload from external captive portal with CoA
type CoARequest struct {
	// required client data
	MAC  string `json:"mac" bson:"mac" form:"mac" query:"mac" validate:"required,mac"`
	CPE  string `json:"cpe_id" bson:"cpe_id" form:"cpe_id" query:"cpe_id" validate:"uuid"`
	WLAN string `json:"wlan_id" bson:"wlan_id" form:"wlan_id" query:"wlan_id" validate:"uuid"`

	// Username will be use in RADIUS Acct
	Username string `json:"username" bson:"username" form:"username" query:"username" validate:"-"`

	// passing NAS-ID to use in next RADIUS Acct
	NasId string `json:"nas-id,omitempty" bson:"nas-id" form:"nas-id" query:"nas-id" validate:"-"`

	// two way to set session-timeout (for backward compatibility)
	Timeout  int64 `json:"session-timeout,omitempty" bson:"session-timeout" form:"session-timeout" query:"session-timeout" validate:"-"`
	Timeout2 int64 `json:"wimark-session-timeout,omitempty" bson:"wimark-session-timeout" validate:"-"`

	// session could be divided on parts
	Transferable bool `json:"transferable"`

	// additional info (from Wimark Portal)

	// Auth additional data
	UserAgent string `json:"useragent"`

	// auth type -- free / sponsor / paid / etc
	AuthType string `json:"auth_type"`

	// authentype -- sms / callback / etc
	AuthenType string `json:"authen_type"`

	// IP address of user -- if has
	IP string `json:"ip" bson:"ip"`

	// Loc ID == location ID
	LocID string `json:"loc-id"`

	// is CNA (Captive Network Assistant)
	CNA bool `json:"cna"`

	// DEPRECATED

	// ACL group -- not using now
	Group string `json:"wimark-client-group,omitempty" bson:"wimark-client-group" validate:"-"`

	Password string `json:"password" bson:"password" form:"password" query:"password" validate:"-"`
}

type CoAResponse struct {
	Status string `json:"status,omitempty"`
	Code   int    `json:"code"`

	Description string      `json:"description,omitempty"`
	Data        interface{} `json:"data,omitempty"`

	State string `json:"state"`
}
