package libwimark

import "time"

// CollClientRF collection for client rf
const CollClientRF = "client_rf"

// ClientRF struct for history client rf data
type ClientRF struct {
	ID       string          `json:"id" bson:"_id"`
	MAC      string          `json:"mac" bson:"mac"`
	CPE      string          `json:"cpe" bson:"cpe"`
	TS       int64           `json:"ts" bson:"ts"`
	CreateAt time.Time       `json:"create_at" bson:"create_at"`
	RF       AccountingRadio `json:"rf" bson:"rf"`
}
