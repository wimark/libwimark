package libwimark

import "time"

const (
	// CollClientRF collection for client rf
	CollClientRF = "client_rf"

	// CollClientDistance collection for client distance from CPE
	CollClientDistance = "client_distance"
)

// ClientRF struct for history client rf data
type ClientRF struct {
	ID       string          `json:"id" bson:"_id,omitempty"`
	MAC      string          `json:"mac" bson:"mac"`
	CPE      string          `json:"cpe" bson:"cpe"`
	Freq     int             `json:"freq" bson:"freq"`
	TS       int64           `json:"ts" bson:"ts"`
	CreateAt time.Time       `json:"create_at" bson:"create_at"`
	RF       AccountingRadio `json:"rf" bson:"rf"`
	SNR      int             `json:"snr" bson:"snr"`
}

// ClientDistance struct for history client distance data
type ClientDistance struct {
	ID       string    `json:"id" bson:"_id,omitempty"`
	MAC      string    `json:"mac" bson:"mac"`
	CPE      string    `json:"cpe" bson:"cpe"`
	Freq     int       `json:"freq" bson:"freq"`
	TS       int64     `json:"ts" bson:"ts"`
	CreateAt time.Time `json:"create_at" bson:"create_at"`
	Signal   int       `json:"signal" bson:"signal"`
	Noise    int       `json:"noise" bson:"noise"`
	SNR      int       `json:"snr" bson:"snr"`
	Distance float64   `json:"distance" bson:"distance"`
}
