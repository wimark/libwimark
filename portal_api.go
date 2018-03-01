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

	Timeout int `json:"timeout,omitempty" bson:"timeout" form:"timeout" query:"timeout" validate:"-"`
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

	MAC  string `json:"mac" bson:"mac"`
	WLAN string `json:"wlan_id" bson:"wlan_id"`
	CPE  string `json:"cpe_id" bson:"cpe_id"`

	AcctFirst int64
	AcctLast  int64
	AuthStart int64
	AuthStop  int64
	Resolved  bool

	RxBytes     int64 `json:"rx_bytes" bson:"rx_bytes"`
	TxBytes     int64 `json:"tx_bytes" bson:"tx_bytes"`
	RxBytesLast int64 `json:"rx_bytes_last" bson:"rx_bytes_last"`
	TxBytesLast int64 `json:"tx_bytes_last" bson:"tx_bytes_last"`
}
