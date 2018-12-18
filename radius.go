package libwimark

// radius.go -- for acct/auth radius packet structure

type RadiusAccessRequest struct {
	Timestamp      int64  `json:"Timestamp"`
	EventTimestamp string `json:"Event-Timestamp"`

	UserName         string `json:"User-Name"`
	Password         string `json:"User-Password"`
	CallingStationId string `json:"Calling-Station-Id"`
	CalledStationId  string `json:"Called-Station-Id"`
	FramedIPAddress  string `json:"Framed-IP-Address"`

	NasPortType   string `json:"NAS-Port-Type"`
	NasPort       string `json:"NAS-Port"`
	NasPortId     string `json:"NAS-Port-Id"`
	NasIdentifier string `json:"NAS-Identifier"`
	NasIPAddress  string `json:"NAS-IP-Address"`
}

type RadiusAccessAccept struct {
}

type RadiusAccessReject struct {
}

type RadiusAccountingRequest struct {
	Timestamp      int64  `json:"Timestamp"`
	EventTimestamp string `json:"Event-Timestamp"`

	UserName         string `json:"User-Name"`
	CallingStationId string `json:"Calling-Station-Id"`
	CalledStationId  string `json:"Called-Station-Id"`
	FramedIPAddress  string `json:"Framed-IP-Address"`

	NasPortType   string `json:"NAS-Port-Type"`
	NasPort       string `json:"NAS-Port"`
	NasPortId     string `json:"NAS-Port-Id"`
	NasIdentifier string `json:"NAS-Identifier"`
	NasIPAddress  string `json:"NAS-IP-Address"`

	AcctStatusType      string `json:"Acct-Status-Type"`
	AcctDelayTime       int    `json:"Acct-Delay-Time"`
	AcctSessionTime     int    `json:"Acct-Session-Time"`
	AcctSessionId       string `json:"Acct-Session-Id"`
	AcctTerminateCause  string `json:"Acct-Terminate-Cause,omitempty"`
	AcctInputGigawords  int    `json:"Acct-Input-Gigawords"`
	AcctOutputGigawords int    `json:"Acct-Output-Gigawords"`
	AcctInputOctets     int    `json:"Acct-Input-Octets"`
	AcctOutputOctets    int    `json:"Acct-Output-Octets"`
	AcctInputPackets    int    `json:"Acct-Input-Packets"`
	AcctOutputPackets   int    `json:"Acct-Output-Packets"`
}
