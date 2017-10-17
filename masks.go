package libwimark

type LimitBetween struct {
	Upper float64 `json:"upper"`
	Lower float64 `json:"lower"`
}

type LimitBetweenOptional struct {
	Upper *float64 `json:"upper"`
	Lower *float64 `json:"lower"`
}

type SimpleMask struct {
	UUID []UUID `json:"uuid"`
}

type TimestampMask struct {
	UUID  []UUID `json:"uuid"`
	Start *int64 `json:"start"`
	Stop  *int64 `json:"stop"`
}

type EventMask struct {
	TimestampMask
	Type       []SystemEventType  `json:"type"`
	Subject_id []string           `json:"subject_id"`
	Level      []SystemEventLevel `json:"level"`
}

type ClientStatMask struct {
	TimestampMask
	CPE              []UUID
	CallingStationId []string
}

type CPEMask struct {
	UUID      []UUID `json:"uuid"`
	HasWLANs  []UUID `json:"has_wlans"`
	Connected *bool  `json:"connected"`
}

type WLANMask struct {
	UUID      []UUID `json:"uuid"`
	HasRadius []UUID `json:"has_radius"`
}

type StatsMask struct {
	UUID    []UUID `json:"uuid"`
	CPEUUID []UUID `json:"cpe"`
	Start   *int64 `json:"start"`
	Stop    *int64 `json:"stop"`
}

type LBSClientDataMask struct {
	TimestampMask
	CPE       []UUID   `json:"cpe"`
	Radio     []string `json:"radio"`
	ClientMac []string `json:"client_mac"`
	RSSI      []int    `json:"rssi"`
}

type LBSCPEInfoMask struct {
	SimpleMask
	Group []UUID               `json:"group"`
	CPE   []UUID               `json:"cpe"`
	Name  []string             `json:"name"`
	X     LimitBetweenOptional `json:"x"`
	Y     LimitBetweenOptional `json:"y"`
	Z     LimitBetweenOptional `json:"z"`
}

type LBSClientCoordsMask struct {
	TimestampMask
	Group []UUID               `json:"group"`
	Mac   []string             `json:"mac"`
	X     LimitBetweenOptional `json:"x"`
	Y     LimitBetweenOptional `json:"y"`
	Z     LimitBetweenOptional `json:"z"`
}

type TunnelMask struct {
	UUID  []UUID `json:"uuid"`
	Hosts []UUID `json:"has_cpes"`
	CPEs  []UUID `json:"has_hosts"`
}
