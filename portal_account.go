package libwimark

import (
	"time"
)

const (
	CollPortalUserAccount = "portal_user_accounts"
	CollPortalUserVoucher = "portal_user_voucher"
	CollPortalPaidPlans   = "portal_paid_plans"
)

// PortalUserAccount struct to represent user account for
// portal
type PortalUserAccount struct {
	ID     string   `json:"id" bson:"_id"`
	MSISDN string   `json:"msisdn" bson:"msisdn"`
	MAC    []string `json:"macs" bson:"macs"`

	Name    string `json:"name" bson:"name"`
	SurName string `json:"surname" bson:"surname"`

	PushAgreement bool `json:"push_agreement" bson:"push_agreement"`

	Balance int `json:"balance" bson:"balance"`

	Create   time.Time `json:"create" bson:"create"`
	CreateAt int64     `json:"create_at" bson:"create_at"`
}

// PortalUserVoucher struct to represent voucher
type PortalUserVoucher struct {
	ID string `json:"id" bson:"_id"`

	Account string `json:"account" bson:"account"`

	Create time.Time `json:"create" bson:"create"`

	CreateAt int64 `json:"create_at" bson:"create_at"`
	ExpireAt int64 `json:"expire_at" bson:"expire_at"`

	PaidPlan struct {
		SpeedDownload int
		SpeedUpload   int
		Amount        int
		Currency      string
	} `json:"tariff" bson:"tariff"`
}

// PortalPaidPlan struct to represent paid plans (tariffs)
type PortalPaidPlan struct {
	SpeedDownload int
	SpeedUpload   int
	Amount        int
	Currency      string
}

type PortalUserBilling struct {
}
