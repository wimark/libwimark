package libwimark

import (
	"time"
)

const (
	CollPortalUserAccount = "portal_user_accounts"
	CollPortalUserVoucher = "portal_user_voucher"
	CollPortalPaidPlans   = "portal_paid_plans"
)

// PortalUserAccount struct to represent user account for WLAN
type PortalUserAccount struct {
	ID string `json:"id" bson:"_id"`

	// WLAN     string   `json:"wlan_id" bson:"wlan_id"`
	Profile  string   `json:"profile" bson:"profile"`
	Identity string   `json:"identity" bson:"identity"`
	MACs     []string `json:"macs" bson:"macs"`

	Name    string `json:"name" bson:"name"`
	SurName string `json:"surname" bson:"surname"`
	Filled  bool   `json:"filled" bson:"filled"`

	PushAgreement bool `json:"push_agreement" bson:"push_agreement"`

	Balance  int    `json:"balance" bson:"balance"`
	Currency string `json:"currency" bson:"currency"`

	Create   time.Time `json:"create" bson:"create"`
	CreateAt int64     `json:"create_at" bson:"create_at"`
}

// PortalUserVoucher struct to represent voucher
type PortalUserVoucher struct {
	ID string `json:"id" bson:"_id"`

	// Profile  string `json:"profile" bson:"profile"`
	// WLAN     string `json:"wlan_id" bson:"wlan_id"`
	// Identity string `json:"identity" bson:"identity"`
	Account string `json:"account" bson:"account"`

	Create   time.Time `json:"create" bson:"create"`
	CreateAt int64     `json:"create_at" bson:"create_at"`

	StartAt  int64 `json:"start_at" bson:"start_at"`
	ExpireAt int64 `json:"expire_at" bson:"expire_at"`

	Plan string `json:"tariff" bson:"tariff"`
}

// PortalTariffPlan struct to represent paid plans (tariffs)
type PortalTariffPlan struct {
	ID string `json:"id" bson:"_id"`

	// name
	Name string `json:"name" bson:"name"`

	// limits
	SpeedLimit   int   `json:"speed" bson:"speed"`
	TimeoutLimit int64 `json:"session" bson:"session"`

	// how much to pay
	Amount   int    `json:"amount" bson:"amount"`
	Currency string `json:"currency" bson:"currency"`

	// service info
	Create   time.Time `json:"create" bson:"create"`
	CreateAt int64     `json:"create_at" bson:"create_at"`
}
