package libwimark

import (
	"time"
)

const (
	COLL_PORTAL_UACCOUNTS   = "portal_user_accounts"
	COLL_PORTAL_UA_VOUCHERS = "portal_user_account_voucher"
)

type PortalUserAccount struct {
	Id string `json:"id" bson:"_id"`

	// user common data
	MAC  string `json:"mac" bson:"mac" form:"mac" query:"mac" validate:"required,mac"`
	WLAN string `json:"wlan_id" bson:"wlan_id" form:"wlan_id" query:"wlan_id" validate:"required,uuid"`

	CreateAt time.Time

	// Ip   string `json:"client_ip" bson:"client_ip" form:"client_ip" query:"client_ip"`
	// CPE  string `json:"cpe_id" bson:"cpe_id" form:"cpe_id" query:"cpe_id" validate:"required,uuid"`

	// AuthenticationData

	// AuthorizationData

	// AccountingData
}
