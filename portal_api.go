package libwimark

import (
	"gopkg.in/go-playground/validator.v9"
)

// Struct for request payload from external captive portal

type RedirectRequestObject struct {
	Username string `json:"username" bson:"username" form:"username" query:"username" validate:"-"`
	Password string `json:"password" bson:"password" form:"password" query:"password" validate:"-"`

	MAC  string `json:"mac" bson:"mac" form:"mac" query:"mac" validate:"required,mac"`
	CPE  string `json:"cpe_id" bson:"cpe_id" form:"cpe_id" query:"cpe_id" validate:"uuid"`
	WLAN string `json:"wlan_id" bson:"wlan_id" form:"wlan_id" query:"wlan_id" validate:"uuid"`

	Timeout int64 `json:"session-timeout,omitempty" bson:"session-timeout" form:"session-timeout" query:"session-timeout" validate:"-"`
}

// Struct for request payload from webui

type PortalRequestObject struct {
	Type string `json:"type,omitempty" bson:"type" form:"type" query:"type" validate:"-"`

	Username string `json:"username,omitempty" bson:"username" form:"username" query:"username" validate:"-"`
	Password string `json:"password,omitempty" bson:"password" form:"password" query:"password" validate:"-"`

	MAC       string `json:"mac" bson:"mac" form:"mac" query:"mac" validate:"required,mac"`
	CPE       string `json:"cpe_id" bson:"cpe_id" form:"cpe_id" query:"cpe_id" validate:"required,uuid"`
	WLAN      string `json:"wlan_id" bson:"wlan_id" form:"wlan_id" query:"wlan_id" validate:"required,uuid"`
	Useragent string `json:"useragent"  bson:"useragent" form:"useragent" query:"useragent" validate:"-"`
	Timeout   int64  `json:"-" validate:"-"`
}

type Validator struct {
	Validator *validator.Validate
}

func (cv *Validator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

type HTTPResponseObject struct {
	Status      string      `json:"status,omitempty"`
	Description string      `json:"description,omitempty"`
	Code        string      `json:"code,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

type RedirectClientSession struct {
	ID string `json:"id" bson:"_id"`

	Session  string `json:"session_id" bson:"session_id"`
	MAC      string `json:"mac" bson:"mac"`
	WLAN     string `json:"wlan_id" bson:"wlan_id"`
	CPE      string `json:"cpe_id" bson:"cpe_id"`
	Redirect string `json:"redirect_id" bson:"redirect_id"`

	AcctStart int64 `json:"acct_start" bson:"acct_start"`
	AuthStart int64 `json:"auth_start" bson:"auth_start"`
	AuthStop  int64 `json:"auth_stop" bson:"auth_stop"`
	Timeout   int64 `json:"timeout" bson:"timeout"`

	// data from acct start
	CallingStationId string `json:"Calling-Station-Id"`
	CalledStationId  string `json:"Called-Station-Id"`
	NasIdentifier    string `json:"NAS-Identifier"`
	NasIPAddress     string `json:"NAS-IP-Address"`
	FramedIPAddress  string `json:"Framed-IP-Address"`
	UserName         string `json:"User-Name"`

	// data for interim update and acct stop
	AcctSessionTime     int `json:"Acct-Session-Time"`
	AcctInputGigawords  int `json:"Acct-Input-Gigawords"`
	AcctOutputGigawords int `json:"Acct-Output-Gigawords"`
	AcctInputOctets     int `json:"Acct-Input-Octets"`
	AcctOutputOctets    int `json:"Acct-Output-Octets"`
	AcctInputPackets    int `json:"Acct-Input-Packets"`
	AcctOutputPackets   int `json:"Acct-Output-Packets"`
}

type PortalAuthObject struct {
	Timestamp int64  `json:"timestamp" bson:"timestamp"`
	Type      string `json:"type" bson:"type"`
	CPE       string `json:"cpe_id" bson:"cpe_id"`
	Useragent string `json:"useragent" bson:"useragent"`
}

type PortalClientSession struct {
	Id string `json:"id" bson:"_id"`

	MAC  string `json:"mac" bson:"mac"`
	WLAN string `json:"wlan_id" bson:"wlan_id"`

	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`

	CreateAt       int64              `json:"create_at" bson:"create_at"`
	StopAt         int64              `json:"stop_at" bson:"stop_at"`
	Timeout        int64              `json:"timeout" bson:"timeout"`
	SessionTimeout int64              `json:"session_timeout" bson:"session_timeout"`
	Auth           []PortalAuthObject `json:"auth" bson:"auth"`
}
