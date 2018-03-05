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

	Timeout int64 `json:"client_timeout,omitempty" bson:"client_timeout" form:"client_timeout" query:"client_timeout" validate:"-"`
}

// Struct for request payload from webui

type PortalRequestObject struct {
	Type string `json:"type,omitempty" bson:"type" form:"type" query:"type" validate:"-"`

	Username string `json:"username,omitempty" bson:"username" form:"username" query:"username" validate:"-"`
	Password string `json:"password,omitempty" bson:"password" form:"password" query:"password" validate:"-"`

	MAC  string `json:"mac" bson:"mac" form:"mac" query:"mac" validate:"required,mac"`
	CPE  string `json:"cpe_id" bson:"cpe_id" form:"cpe_id" query:"cpe_id" validate:"required,uuid"`
	WLAN string `json:"wlan_id" bson:"wlan_id" form:"wlan_id" query:"wlan_id" validate:"required,uuid"`
}

type Validator struct {
	Validator *validator.Validate
}

func (cv *Validator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

type HTTPResponseObject struct {
	Status      string `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
	Code        string `json:"code,omitempty"`
}

type PortalClientSession struct {
	ID string `json:"id" bson:"_id"`

	MAC      string `json:"mac" bson:"mac"`
	WLAN     string `json:"wlan_id" bson:"wlan_id"`
	CPE      string `json:"cpe_id" bson:"cpe_id"`
	Redirect string `json:"redirect_id" bson:"redirect_id"`
	Session  string `json:"session_id" bson:"session_id"`

	AcctStart int64 `json:"acct_start" bson:"acct_start"`
	AcctStop  int64 `json:"acct_stop" bson:"acct_stop"`
	AuthStart int64 `json:"auth_start" bson:"auth_start"`
	AuthStop  int64 `json:"auth_stop" bson:"auth_stop"`
	Timeout   int64 `json:"timeout" bson:"timeout"`

	RxBytes     int64 `json:"rx_bytes" bson:"rx_bytes"`
	TxBytes     int64 `json:"tx_bytes" bson:"tx_bytes"`
	RxBytesLast int64 `json:"rx_bytes_last" bson:"rx_bytes_last"`
	TxBytesLast int64 `json:"tx_bytes_last" bson:"tx_bytes_last"`
}
