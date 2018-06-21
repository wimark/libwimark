package libwimark

// radius.go -- for acct/auth radius packet structure

type RadiusAccessRequest struct {
}

type RadiusAccessAccept struct {
}

type RadiusAccessReject struct {
}

type RadiusAccountingRequest struct {
	//        Acct-Status-Type = Start
	//        NAS-Port-Type = Wireless-802.11
	//        Calling-Station-Id = "40:D3:AE:04:32:0C"
	//        Called-Station-Id = "hotspot1"
	//        NAS-Port-Id = "HotSpot"
	//        User-Name = "40:D3:AE:04:32:0C"
	//        NAS-Port = 2150629507
	//        Acct-Session-Id = "80300083"
	//        Framed-IP-Address = 192.168.2.11
	//        Mikrotik-Host-IP = 192.168.2.11
	//        Event-Timestamp = "Mar  5 2018 16:19:46 GMT"
	//        NAS-Identifier = "8t37nd4K5ekznE7a"
	//        Acct-Delay-Time = 0
	//        NAS-IP-Address = 10.231.0.214
	//        Acct-Unique-Session-Id = "ce4df8c7d20d469b"
	//        Timestamp = 1520266785

	AcctStatusType     string `json:"Acct-Status-Type"`
	CallingStationId   string `json:"Calling-Station-Id"`
	CalledStationId    string `json:"Called-Station-Id"`
	FramedIPAddress    string `json:"Framed-IP-Address"`
	UserName           string `json:"User-Name"`
	Timestamp          int64  `json:"Timestamp"`
	EventTimestamp     string `json:"Event-Timestamp"`
	NasPortType        string `json:"NAS-Port-Type"`
	NasPort            string `json:"NAS-Port"`
	NasIdentifier      string `json:"NAS-Identifier"`
	NasIPAddress       string `json:"NAS-IP-Address"`
	AcctDelayTime      int    `json:"Acct-Delay-Time"`
	AcctSessionTime    int    `json:"Acct-Session-Time"`
	AcctSessionId      string `json:"Acct-Session-Id"`
	AcctTerminateCause string `json:"Acct-Terminate-Cause,omitempty"`

	AcctInputGigawords  int `json:"Acct-Input-Gigawords"`
	AcctOutputGigawords int `json:"Acct-Output-Gigawords"`
	AcctInputOctets     int `json:"Acct-Input-Octets"`
	AcctOutputOctets    int `json:"Acct-Output-Octets"`
	AcctInputPackets    int `json:"Acct-Input-Packets"`
	AcctOutputPackets   int `json:"Acct-Output-Packets"`
}
