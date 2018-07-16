package libwimark

type BackScanModelRaw struct {
	SSID        string          `json:"ssid" bson:"ssid"`
	BSSID       string          `json:"bssid" bson:"bssid"`
	Signal      float32         `json:"signal" bson:"signal"`
	Channel     int             `json:"channel" bson:"channel"`
	Quality     float32         `json:"quality" bson:"quality"`
	QualityMax  int             `json:"quality_max" bson:"quality_max"`
	Mode        string          `json:"mode" bson:"mode"`
	CenterIndex int             `json:"center_idx0" bson:"center_idx0"`
	ChanOffset  int             `json:"sec_channel_offset" bson:"sec_channel_offset"`
	HTModes     map[string]bool `json:"htmodelist" bson:"htmodelist"`
	Encryption  struct {
		Enabled      bool     `json:"enabled" bson:"enabled"`
		AuthAlgs     []string `json:"auth_algs" bson:"auth_algs"`
		Description  string   `json:"description" bson:"description"`
		WEP          bool     `json:"wep" bson:"wep"`
		WPA          int      `json:"wpa" bson:"wpa"`
		AuthSuites   []string `json:"auth_suites" bson:"auth_suites"`
		PairCiphers  []string `json:"pair_ciphers" bson:"pair_ciphers"`
		GroupCiphers []string `json:"group_ciphers" bson:"group_ciphers"`
	} `json:"encryption" bson:"encryption"`
}

type CPEScanData struct {
	CPE       UUID                        `json:"cpe" bson:"cpe"`
	Radio     string                      `json:"radio" bson:"radio"`
	Timestamp int64                       `json:"timestamp" bson:"timestamp"`
	Scanlist  map[string]BackScanModelRaw `json:"scanlist" bson:"scanlist"`
}
