package libwimark

// SNMPGeneral -
type SNMPGeneral struct {
	SystemLocation string `json:"system_location" bson:"system_location"`
	SystemContacts string `json:"system_contacts" bson:"system_contacts"`
	Traps          struct {
		Available []string `json:"available" bson:"available"`
		Enabled   []string `json:"enabled" bson:"enabled"`
	} `json:"traps" bson:"traps"`

	SNMPView          `json:"snmp_view" bson:"snmp_view"`
	SNMPWirelessTraps `json:"snmp_wireless_traps"`
}

type SNMPView struct {
	Included []string `json:"included" bson:"included"`
	Excluded []string `json:"excluded" bson:"excluded"`
}

// SNMPCommunityString -
type SNMPCommunityString struct {
	ID            UUID           `json:"id" bson:"_id"`
	CommunityName string         `json:"community_name" bson:"community_name"`
	AccessMode    SNMPAccessMode `json:"access_mode" bson:"access_mode"`
}

// SNMPV3Users -
type SNMPV3Users struct {
	ID                UUID                  `json:"id" bson:"_id"`
	UserName          string                `json:"user_name" bson:"user_name"`
	UserGroup         UUID                  `json:"user_group" bson:"user_group"`
	SecurityLevel     SNMPSecurityLevelType `json:"security_level" bson:"security_level"`
	AuthProtocol      SNMPAuthProtocol      `json:"auth_protocol" bson:"auth_protocol"`
	AuthPassPhrase    string                `json:"auth_pass_phrase" bson:"auth_pass_phrase"`
	PrivacyProtocol   SNMPPrivacyProtocol   `json:"privacy_protocol" bson:"privacy_protocol"`
	PrivacyPassPhrase string                `json:"privacy_pass_phrase" bson:"privacy_pass_phrase"`
}

// SNMPV3UserGroups -
type SNMPV3UserGroups struct {
	ID            UUID                  `json:"id" bson:"_id"`
	Name          string                `json:"name" bson:"name"`
	SecurityLevel SNMPSecurityLevelType `json:"security_level" bson:"security_level"`
	ReadView      string                `json:"read_view" bson:"read_view"`
	WriteView     string                `json:"write_view" bson:"write_view"`
}

// SNMPHost -
type SNMPHost struct {
	ID              UUID        `json:"id" bson:"_id"`
	IPAddress       string      `json:"ip_address" bson:"ip_address"`
	Version         SNMPVersion `json:"version" bson:"version"`
	Port            int         `json:"port" bson:"port"`
	CommunityString UUID        `json:"community_string" bson:"community_string"`
	Type            string      `json:"type" bson:"type"`
}

// SNMPWirelessTraps -
type SNMPWirelessTraps struct {
	MeshTrap             `json:"mesh_trap" bson:"mesh_trap"`
	RFTrap               `json:"rf_trap" bson:"rf_trap"`
	WirelessMobilityTrap `json:"wireless_mobility_trap" bson:"wireless_mobility_trap"`
	RRMTrap              `json:"rrm_trap" bson:"rrm_trap"`
	WirelessClientTrap   `json:"wireless_client_trap" bson:"wireless_client_trap"`
	RoqueTrap            `json:"roque_trap" bson:"roque_trap"`
	APTrap               `json:"ap_trap" bson:"ap_trap"`
	GeneralController    bool `json:"general_controller" bson:"general_controller"`
}

// MeshTrap -
type MeshTrap struct {
	AbateSNR          bool `json:"abate_snr" bson:"abate_snr"`
	ChildMoved        bool `json:"child_moved:" bson:"child_moved"`
	ExcessiveHopCount bool `json:"excessive_hop_count" bson:"excessive_hop_count"`
	ParentChange      bool `json:"parent_change" bson:"parent_change"`
	AuthFailure       bool `json:"auth_failure" bson:"auth_failure"`
	ExcessiveChildren bool `json:"excessive_children" bson:"excessive_children"`
	OnsetSNR          bool `json:"onset_snr" bson:"onset_snr"`
}

// RFTrap -
type RFTrap struct {
	TxPower      bool `json:"tx_power" bson:"tx_power"`
	Coverage     bool `json:"coverage" bson:"coverage"`
	Load         bool `json:"load" bson:"load"`
	Channels     bool `json:"channels" bson:"channels"`
	Interference bool `json:"interference" bson:"interference"`
	Noise        bool `json:"noise" bson:"noise"`
}

// WirelessMobilityTrap -
type WirelessMobilityTrap struct {
	AnchorClientEnabled bool `json:"anchor_client_enabled" bson:"anchor_client_enabled"`
}

// RRMTrap -
type RRMTrap struct {
	Group bool `json:"group" bson:"group"`
}

// WirelessClientTrap -
type WirelessClientTrap struct {
	AssociationFail     bool `json:"association_fail" bson:"association_fail"`
	AuthenticationFail  bool `json:"authentication_fail" bson:"authentication_fail"`
	Disassociate        bool `json:"disassociate" bson:"disassociate"`
	Authenticate        bool `json:"authenticate" bson:"authenticate"`
	Associate           bool `json:"associate" bson:"associate"`
	DeAuthenticate      bool `json:"de_authenticate" bson:"de_authenticate"`
	Excluded            bool `json:"excluded" bson:"excluded"`
	MaxThresholdWarning bool `json:"max_threshold_warning" bson:"maxThresholdWarning"`
}

// RoqueTrap -
type RoqueTrap struct {
	RogueClient bool `json:"rogue_client" bson:"rogue_client"`
	RogueAP     bool `json:"rogue_ap" bson:"rogue_ap"`
}

// APTrap -
type APTrap struct {
	Crash                   bool `json:"crash" bson:"crash"`
	IPAddressFallBack       bool `json:"ip_address_fall_back" bson:"ip_address_fall_back"`
	NoRadioCards            bool `json:"no_radio_cards" bson:"no_radio_cards"`
	MFP                     bool `json:"mfp" bson:"mfp"`
	DualBandRadioRoleChange bool `json:"dual_band_radio_role_change" bson:"dual_band_radio_role_change"`
	Authorization           bool `json:"authorization" bson:"authorization"`
	InterfaceUp             bool `json:"interface_up" bson:"interface_up"`
	Mode                    bool `json:"mode" bson:"mode"`
	Register                bool `json:"register" bson:"register"`
	BrokenAntenna           bool `json:"broken_antenna" bson:"broken_antenna"`
	DualBandRadioBandChange bool `json:"dual_band_radio_band_change" bson:"dual_band_radio_band_change"`
	APStats                 bool `json:"ap_stats" bson:"ap_stats"`
}
