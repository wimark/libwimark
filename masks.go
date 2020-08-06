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
	UUID                []UUID   `json:"uuid"`
	HasWLANs            []UUID   `json:"has_wlans"`
	Connected           *bool    `json:"connected"`
	HasL2Chains         []UUID   `json:"has_l2chains"`
	HasController       []string `json:"has_controller"`
	HasCaptiveRedirects []UUID   `json:"has_redirects"`
}

type WLANMask struct {
	UUID                []UUID `json:"uuid"`
	HasRadius           []UUID `json:"has_radius"`
	HasL2Chains         []UUID `json:"has_l2chains"`
	HasCaptiveRedirects []UUID `json:"has_redirects"`
	HasHotspotProfiles  []UUID `json:"has_hotspots"`
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

type LBSZoneMask struct {
	SimpleMask
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Group       UUID               `json:"group"`
	Corners     []CornerCoordsMask `json:"corners"`
}

type CornerCoordsMask struct {
	SimpleMask
	X LimitBetweenOptional `json:"x" bson:"x"`
	Y LimitBetweenOptional `json:"y" bson:"y"`
	Z LimitBetweenOptional `json:"z" bson:"z"`
}

type LBSClientCoordsMask struct {
	TimestampMask
	Group []UUID               `json:"group"`
	Mac   []string             `json:"mac"`
	X     LimitBetweenOptional `json:"x"`
	Y     LimitBetweenOptional `json:"y"`
	Z     LimitBetweenOptional `json:"z"`
}

type CPEModelMask struct {
	UUID  []UUID   `json:"uuid"`
	Names []string `json:"names"`
}

type ConfigRuleMask struct {
	UUID        []UUID `json:"uuid"`
	CPEs        []UUID `json:"has_cpes"`
	Models      []UUID `json:"has_models"`
	WLANs       []UUID `json:"has_wlans"`
	HasL2Chains []UUID `json:"has_l2chains"`
	Auto        *bool  `json:"is_auto"`
	Always      *bool  `json:"is_always"`
}

type ControllerMask struct {
	UUID    []UUID `json:"uuid"`
	Enabled *bool  `json:"is_enabled"`
}
