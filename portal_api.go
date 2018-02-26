package libwimark

import (
	"gopkg.in/go-playground/validator.v9"
)

// Struct for request payload from external captive portal

type RedirectRequestObject struct {
	Username string `json:"username" bson:"username" form:"username" query:"username" validate:"required"`
	Password string `json:"password" bson:"password" form:"password" query:"password" validate:"required"`
	MAC      string `json:"mac" bson:"mac" form:"mac" query:"mac" validate:"required,len=17"`
	//Portal   string `json:"portal_id" bson:"portal_id" form:"portal_id" query:"portal_id" validate:"required"`

	CPE  string `json:"cpe_id" bson:"cpe_id" form:"cpe_id" query:"cpe_id" validate:"required"`
	WLAN string `json:"wlan_id" bson:"wlan_id" form:"wlan_id" query:"wlan_id" validate:"required"`
}

// Struct for request payload from webui

type PortalRequestObject struct {
	Type string `json:"type,omitempty" bson:"type" form:"type" query:"type" validate:"-"`

	Username string `json:"username,omitempty" bson:"username" form:"username" query:"username" validate:"-"`
	Password string `json:"password,omitempty" bson:"password" form:"password" query:"password" validate:"-"`

	MAC  string `json:"mac" bson:"mac" form:"mac" query:"mac" validate:"required"`
	CPE  string `json:"cpe_id" bson:"cpe_id" form:"cpe_id" query:"cpe_id" validate:"required"`
	WLAN string `json:"wlan_id" bson:"wlan_id" form:"wlan_id" query:"wlan_id" validate:"required"`
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
}
