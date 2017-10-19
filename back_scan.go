package libwimark

type BackScanModelRaw struct {
	SSID       string `json:"ssid" bson:"ssid"`
	BSSID      string `json:"bssid" bson:"bssid"`
	Signal     int    `json:"signal" bson:"signal"`
	Channel    int    `json:"channel" bson:"channel"`
	Quality    int    `json:"quality" bson:"quality"`
	QualityMax int    `json:"quality_max" bson:"quality_max"`
	Mode       string `json:"mode" bson:"mode"`
	Encryption struct {
		Enabled      bool     `json:"enabled" bson:"enabled"`
		AuthAlgs     []string `json:"auth_algs" bson:"auth_algs"`
		Description  string   `json:"description" bson:"description"`
		WEP          bool     `json:"" bson:""`
		WPA          int      `json:"wpa" bson:"wep"`
		AuthSuites   []string `json:"auth_suites" bson:"auth_suites"`
		PairCiphers  []string `json:"pair_ciphers" bson:"pair_ciphers"`
		GroupCiphers []string `json:"group_ciphers" bson:"group_ciphers"`
	} `json:"encryption" bson:"encryption"`
}
