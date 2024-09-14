package ec2

import types "libds/aws/types"

type VpnConnection struct {
	// The CIDR block of the inside IP addresses for the first VPN tunnel. Valid value is a size /30 CIDR block from the 169.254.0.0/16 range.
	Tunnel1InsideCidr string `json:"tunnel1InsideCidr,omitempty" yaml:"tunnel1InsideCidr,omitempty"`

	// List of one or more integrity algorithms that are permitted for the second VPN tunnel for phase 2 IKE negotiations. Valid values are `SHA1 | SHA2-256 | SHA2-384 | SHA2-512`.
	Tunnel2Phase2IntegrityAlgorithms []string `json:"tunnel2Phase2IntegrityAlgorithms,omitempty" yaml:"tunnel2Phase2IntegrityAlgorithms,omitempty"`

	// The lifetime for phase 2 of the IKE negotiation for the second VPN tunnel, in seconds. Valid value is between `900` and `3600`.
	Tunnel2Phase2LifetimeSeconds int `json:"tunnel2Phase2LifetimeSeconds,omitempty" yaml:"tunnel2Phase2LifetimeSeconds,omitempty"`

	// List of one or more integrity algorithms that are permitted for the first VPN tunnel for phase 2 IKE negotiations. Valid values are `SHA1 | SHA2-256 | SHA2-384 | SHA2-512`.
	Tunnel1Phase2IntegrityAlgorithms []string `json:"tunnel1Phase2IntegrityAlgorithms,omitempty" yaml:"tunnel1Phase2IntegrityAlgorithms,omitempty"`

	// The preshared key of the first VPN tunnel. The preshared key must be between 8 and 64 characters in length and cannot start with zero(0). Allowed characters are alphanumeric characters, periods(.) and underscores(_).
	Tunnel1PresharedKey string `json:"tunnel1PresharedKey,omitempty" yaml:"tunnel1PresharedKey,omitempty"`

	// The IKE versions that are permitted for the second VPN tunnel. Valid values are `ikev1 | ikev2`.
	Tunnel2IkeVersions []string `json:"tunnel2IkeVersions,omitempty" yaml:"tunnel2IkeVersions,omitempty"`

	// The lifetime for phase 1 of the IKE negotiation for the second VPN tunnel, in seconds. Valid value is between `900` and `28800`.
	Tunnel2Phase1LifetimeSeconds int `json:"tunnel2Phase1LifetimeSeconds,omitempty" yaml:"tunnel2Phase1LifetimeSeconds,omitempty"`

	// The percentage of the rekey window for the first VPN tunnel (determined by `tunnel1_rekey_margin_time_seconds`) during which the rekey time is randomly selected. Valid value is between `0` and `100`.
	Tunnel1RekeyFuzzPercentage int `json:"tunnel1RekeyFuzzPercentage,omitempty" yaml:"tunnel1RekeyFuzzPercentage,omitempty"`

	// The IPv6 CIDR on the customer gateway (on-premises) side of the VPN connection.
	LocalIpv6NetworkCidr string `json:"localIpv6NetworkCidr,omitempty" yaml:"localIpv6NetworkCidr,omitempty"`

	// Options for logging VPN tunnel activity. See Log Options below for more details.
	Tunnel1LogOptions types.Ec2_VpnConnectionTunnel1LogOptions `json:"tunnel1LogOptions,omitempty" yaml:"tunnel1LogOptions,omitempty"`

	// One or more integrity algorithms that are permitted for the first VPN tunnel for phase 1 IKE negotiations. Valid values are `SHA1 | SHA2-256 | SHA2-384 | SHA2-512`.
	Tunnel1Phase1IntegrityAlgorithms []string `json:"tunnel1Phase1IntegrityAlgorithms,omitempty" yaml:"tunnel1Phase1IntegrityAlgorithms,omitempty"`

	// The margin time, in seconds, before the phase 2 lifetime expires, during which the AWS side of the first VPN connection performs an IKE rekey. The exact time of the rekey is randomly selected based on the value for `tunnel1_rekey_fuzz_percentage`. Valid value is between `60` and half of `tunnel1_phase2_lifetime_seconds`.
	Tunnel1RekeyMarginTimeSeconds int `json:"tunnel1RekeyMarginTimeSeconds,omitempty" yaml:"tunnel1RekeyMarginTimeSeconds,omitempty"`

	// List of one or more encryption algorithms that are permitted for the second VPN tunnel for phase 2 IKE negotiations. Valid values are `AES128 | AES256 | AES128-GCM-16 | AES256-GCM-16`.
	Tunnel2Phase2EncryptionAlgorithms []string `json:"tunnel2Phase2EncryptionAlgorithms,omitempty" yaml:"tunnel2Phase2EncryptionAlgorithms,omitempty"`

	// The IPv4 CIDR on the AWS side of the VPN connection.
	RemoteIpv4NetworkCidr string `json:"remoteIpv4NetworkCidr,omitempty" yaml:"remoteIpv4NetworkCidr,omitempty"`

	// The action to take after DPD timeout occurs for the first VPN tunnel. Specify restart to restart the IKE initiation. Specify clear to end the IKE session. Valid values are `clear | none | restart`.
	Tunnel1DpdTimeoutAction string `json:"tunnel1DpdTimeoutAction,omitempty" yaml:"tunnel1DpdTimeoutAction,omitempty"`

	// The number of packets in an IKE replay window for the first VPN tunnel. Valid value is between `64` and `2048`.
	Tunnel1ReplayWindowSize int `json:"tunnel1ReplayWindowSize,omitempty" yaml:"tunnel1ReplayWindowSize,omitempty"`

	// List of one or more encryption algorithms that are permitted for the first VPN tunnel for phase 1 IKE negotiations. Valid values are `AES128 | AES256 | AES128-GCM-16 | AES256-GCM-16`.
	Tunnel1Phase1EncryptionAlgorithms []string `json:"tunnel1Phase1EncryptionAlgorithms,omitempty" yaml:"tunnel1Phase1EncryptionAlgorithms,omitempty"`

	// The number of seconds after which a DPD timeout occurs for the second VPN tunnel. Valid value is equal or higher than `30`.
	Tunnel2DpdTimeoutSeconds int `json:"tunnel2DpdTimeoutSeconds,omitempty" yaml:"tunnel2DpdTimeoutSeconds,omitempty"`

	// The action to take after DPD timeout occurs for the second VPN tunnel. Specify restart to restart the IKE initiation. Specify clear to end the IKE session. Valid values are `clear | none | restart`.
	Tunnel2DpdTimeoutAction string `json:"tunnel2DpdTimeoutAction,omitempty" yaml:"tunnel2DpdTimeoutAction,omitempty"`

	// Indicates if a Public S2S VPN or Private S2S VPN over AWS Direct Connect. Valid values are `PublicIpv4 | PrivateIpv4`
	OutsideIpAddressType string `json:"outsideIpAddressType,omitempty" yaml:"outsideIpAddressType,omitempty"`

	// . The attachment ID of the Transit Gateway attachment to Direct Connect Gateway. The ID is obtained through a data source only.
	TransportTransitGatewayAttachmentId string `json:"transportTransitGatewayAttachmentId,omitempty" yaml:"transportTransitGatewayAttachmentId,omitempty"`

	// The action to take when the establishing the tunnel for the first VPN connection. By default, your customer gateway device must initiate the IKE negotiation and bring up the tunnel. Specify start for AWS to initiate the IKE negotiation. Valid values are `add | start`.
	Tunnel1StartupAction string `json:"tunnel1StartupAction,omitempty" yaml:"tunnel1StartupAction,omitempty"`

	// Indicate whether to enable acceleration for the VPN connection. Supports only EC2 Transit Gateway.
	EnableAcceleration bool `json:"enableAcceleration,omitempty" yaml:"enableAcceleration,omitempty"`

	// The IPv4 CIDR on the customer gateway (on-premises) side of the VPN connection.
	LocalIpv4NetworkCidr string `json:"localIpv4NetworkCidr,omitempty" yaml:"localIpv4NetworkCidr,omitempty"`

	// Indicate whether the VPN tunnels process IPv4 or IPv6 traffic. Valid values are `ipv4 | ipv6`. `ipv6` Supports only EC2 Transit Gateway.
	TunnelInsideIpVersion string `json:"tunnelInsideIpVersion,omitempty" yaml:"tunnelInsideIpVersion,omitempty"`

	// Whether the VPN connection uses static routes exclusively. Static routes must be used for devices that don't support BGP.
	StaticRoutesOnly bool `json:"staticRoutesOnly,omitempty" yaml:"staticRoutesOnly,omitempty"`

	// The IKE versions that are permitted for the first VPN tunnel. Valid values are `ikev1 | ikev2`.
	Tunnel1IkeVersions []string `json:"tunnel1IkeVersions,omitempty" yaml:"tunnel1IkeVersions,omitempty"`

	// List of one or more encryption algorithms that are permitted for the second VPN tunnel for phase 1 IKE negotiations. Valid values are `AES128 | AES256 | AES128-GCM-16 | AES256-GCM-16`.
	Tunnel2Phase1EncryptionAlgorithms []string `json:"tunnel2Phase1EncryptionAlgorithms,omitempty" yaml:"tunnel2Phase1EncryptionAlgorithms,omitempty"`

	// The number of packets in an IKE replay window for the second VPN tunnel. Valid value is between `64` and `2048`.
	Tunnel2ReplayWindowSize int `json:"tunnel2ReplayWindowSize,omitempty" yaml:"tunnel2ReplayWindowSize,omitempty"`

	// List of one or more Diffie-Hellman group numbers that are permitted for the second VPN tunnel for phase 2 IKE negotiations. Valid values are `2 | 5 | 14 | 15 | 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23 | 24`.
	Tunnel2Phase2DhGroupNumbers []int `json:"tunnel2Phase2DhGroupNumbers,omitempty" yaml:"tunnel2Phase2DhGroupNumbers,omitempty"`

	// The lifetime for phase 2 of the IKE negotiation for the first VPN tunnel, in seconds. Valid value is between `900` and `3600`.
	Tunnel1Phase2LifetimeSeconds int `json:"tunnel1Phase2LifetimeSeconds,omitempty" yaml:"tunnel1Phase2LifetimeSeconds,omitempty"`

	// The range of inside IPv6 addresses for the second VPN tunnel. Supports only EC2 Transit Gateway. Valid value is a size /126 CIDR block from the local fd00::/8 range.
	Tunnel2InsideIpv6Cidr string `json:"tunnel2InsideIpv6Cidr,omitempty" yaml:"tunnel2InsideIpv6Cidr,omitempty"`

	// The preshared key of the second VPN tunnel. The preshared key must be between 8 and 64 characters in length and cannot start with zero(0). Allowed characters are alphanumeric characters, periods(.) and underscores(_).
	Tunnel2PresharedKey string `json:"tunnel2PresharedKey,omitempty" yaml:"tunnel2PresharedKey,omitempty"`

	// The ID of the EC2 Transit Gateway.
	TransitGatewayId string `json:"transitGatewayId,omitempty" yaml:"transitGatewayId,omitempty"`

	// The number of seconds after which a DPD timeout occurs for the first VPN tunnel. Valid value is equal or higher than `30`.
	Tunnel1DpdTimeoutSeconds int `json:"tunnel1DpdTimeoutSeconds,omitempty" yaml:"tunnel1DpdTimeoutSeconds,omitempty"`

	// The percentage of the rekey window for the second VPN tunnel (determined by `tunnel2_rekey_margin_time_seconds`) during which the rekey time is randomly selected. Valid value is between `0` and `100`.
	Tunnel2RekeyFuzzPercentage int `json:"tunnel2RekeyFuzzPercentage,omitempty" yaml:"tunnel2RekeyFuzzPercentage,omitempty"`

	// The type of VPN connection. The only type AWS supports at this time is "ipsec.1".
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The range of inside IPv6 addresses for the first VPN tunnel. Supports only EC2 Transit Gateway. Valid value is a size /126 CIDR block from the local fd00::/8 range.
	Tunnel1InsideIpv6Cidr string `json:"tunnel1InsideIpv6Cidr,omitempty" yaml:"tunnel1InsideIpv6Cidr,omitempty"`

	// Tags to apply to the connection. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Turn on or off tunnel endpoint lifecycle control feature for the first VPN tunnel. Valid values are `true | false`.
	Tunnel1EnableTunnelLifecycleControl bool `json:"tunnel1EnableTunnelLifecycleControl,omitempty" yaml:"tunnel1EnableTunnelLifecycleControl,omitempty"`

	// The action to take when the establishing the tunnel for the second VPN connection. By default, your customer gateway device must initiate the IKE negotiation and bring up the tunnel. Specify start for AWS to initiate the IKE negotiation. Valid values are `add | start`.
	Tunnel2StartupAction string `json:"tunnel2StartupAction,omitempty" yaml:"tunnel2StartupAction,omitempty"`

	// List of one or more Diffie-Hellman group numbers that are permitted for the first VPN tunnel for phase 1 IKE negotiations. Valid values are ` 2 | 14 | 15 | 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23 | 24`.
	Tunnel1Phase1DhGroupNumbers []int `json:"tunnel1Phase1DhGroupNumbers,omitempty" yaml:"tunnel1Phase1DhGroupNumbers,omitempty"`

	// The lifetime for phase 1 of the IKE negotiation for the first VPN tunnel, in seconds. Valid value is between `900` and `28800`.
	Tunnel1Phase1LifetimeSeconds int `json:"tunnel1Phase1LifetimeSeconds,omitempty" yaml:"tunnel1Phase1LifetimeSeconds,omitempty"`

	// The CIDR block of the inside IP addresses for the second VPN tunnel. Valid value is a size /30 CIDR block from the 169.254.0.0/16 range.
	Tunnel2InsideCidr string `json:"tunnel2InsideCidr,omitempty" yaml:"tunnel2InsideCidr,omitempty"`

	// Options for logging VPN tunnel activity. See Log Options below for more details.
	Tunnel2LogOptions types.Ec2_VpnConnectionTunnel2LogOptions `json:"tunnel2LogOptions,omitempty" yaml:"tunnel2LogOptions,omitempty"`

	// Turn on or off tunnel endpoint lifecycle control feature for the second VPN tunnel. Valid values are `true | false`.
	Tunnel2EnableTunnelLifecycleControl bool `json:"tunnel2EnableTunnelLifecycleControl,omitempty" yaml:"tunnel2EnableTunnelLifecycleControl,omitempty"`

	// The ID of the customer gateway.
	CustomerGatewayId string `json:"customerGatewayId,omitempty" yaml:"customerGatewayId,omitempty"`

	// The IPv6 CIDR on the AWS side of the VPN connection.
	RemoteIpv6NetworkCidr string `json:"remoteIpv6NetworkCidr,omitempty" yaml:"remoteIpv6NetworkCidr,omitempty"`

	// List of one or more encryption algorithms that are permitted for the first VPN tunnel for phase 2 IKE negotiations. Valid values are `AES128 | AES256 | AES128-GCM-16 | AES256-GCM-16`.
	Tunnel1Phase2EncryptionAlgorithms []string `json:"tunnel1Phase2EncryptionAlgorithms,omitempty" yaml:"tunnel1Phase2EncryptionAlgorithms,omitempty"`

	// The margin time, in seconds, before the phase 2 lifetime expires, during which the AWS side of the second VPN connection performs an IKE rekey. The exact time of the rekey is randomly selected based on the value for `tunnel2_rekey_fuzz_percentage`. Valid value is between `60` and half of `tunnel2_phase2_lifetime_seconds`.
	Tunnel2RekeyMarginTimeSeconds int `json:"tunnel2RekeyMarginTimeSeconds,omitempty" yaml:"tunnel2RekeyMarginTimeSeconds,omitempty"`

	// List of one or more Diffie-Hellman group numbers that are permitted for the first VPN tunnel for phase 2 IKE negotiations. Valid values are `2 | 5 | 14 | 15 | 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23 | 24`.
	Tunnel1Phase2DhGroupNumbers []int `json:"tunnel1Phase2DhGroupNumbers,omitempty" yaml:"tunnel1Phase2DhGroupNumbers,omitempty"`

	// List of one or more Diffie-Hellman group numbers that are permitted for the second VPN tunnel for phase 1 IKE negotiations. Valid values are ` 2 | 14 | 15 | 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23 | 24`.
	Tunnel2Phase1DhGroupNumbers []int `json:"tunnel2Phase1DhGroupNumbers,omitempty" yaml:"tunnel2Phase1DhGroupNumbers,omitempty"`

	// One or more integrity algorithms that are permitted for the second VPN tunnel for phase 1 IKE negotiations. Valid values are `SHA1 | SHA2-256 | SHA2-384 | SHA2-512`.
	Tunnel2Phase1IntegrityAlgorithms []string `json:"tunnel2Phase1IntegrityAlgorithms,omitempty" yaml:"tunnel2Phase1IntegrityAlgorithms,omitempty"`

	// The ID of the Virtual Private Gateway.
	VpnGatewayId string `json:"vpnGatewayId,omitempty" yaml:"vpnGatewayId,omitempty"`
}
