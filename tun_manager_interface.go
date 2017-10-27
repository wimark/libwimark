package libwimark

// Tun Manager

type TunManagerBroadcastMeta struct {
	Hostname       string                 `json:"hostname"`
	HostUUID       string                 `json:"host_uuid"`
	HostIP         string                 `json:"host_ip"`
	HostInterfaces []LinkDescriptor       `json:"active_out_interfaces"`
	HostTunnels    []CPETunnelDescription `json:"active_cpe_tunnels"`
}

type LinkDescriptor struct {
	LinkName   string `json:"name"`
	LinkType   string `json:"type"`
	LinkHWAddr string `json:"hw_addr"`
}

type CPETunnelDescription struct {
	TunnelID           int      `json:"tunnel_id"`
	InterfacesAttached []string `json:"interfaces_attached"`
}

// // CreateL2Tunnel

// type CreateL2TunnelParams struct {
// 	CPETunnelID int    `json:"cpe_tunnel_id"`
// 	CPEIP       string `json:"cpe_ip"`
// }

// type CreateL2TunnelResult struct {
// 	WasCreated   bool `json:"was_created"`
// 	HostTunnelID int  `json:"host_tunnel_id"`
// }

// // DeleteL2Tunnel

// type DeleteL2TunnelParams struct {
// 	HostTunnelID int `json:"host_tunnel_id"`
// }

// type DeleteL2TunnelResult struct {
// 	WasDeleted bool `json:"was_deleted"`
// }

// CreateL2TunnelSession

type CreateL2TunnelSessionParams struct {
	CPEIP                    string `json:"cpe_ip"`
	CPESessionID             int    `json:"cpe_session_id"`
	CPETunnelID              int    `json:"cpe_tunnel_id"`
	CommutationInterfaceName string `json:"commutation_interface"`
}

type CreateL2TunnelSessionResult struct {
	WasCreated          bool   `json:"was_created"`
	HostSessionID       int    `json:"host_session_id"`
	HostTunnelID        int    `json:"host_tunnel_id"`
	HostL2InterfaceName string `json:"host_l2_interface"`
}

// DeleteL2TunnelSession

type DeleteL2TunnelSessionParams struct {
	HostSessionID     int    `json:"host_session_id"`
	HostInterfaceName string `json:"host_l2_interface"`
}

type DeleteL2TunnelSessionResult struct {
	WasDeleted       bool `json:"was_deleted"`
	WasTunnelDeleted bool `json:"was_tunnel_deleted"`
}
