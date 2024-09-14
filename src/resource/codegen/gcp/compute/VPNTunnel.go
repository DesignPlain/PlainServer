package compute

type VPNTunnel struct {
	/*
	   Local traffic selector to use when establishing the VPN tunnel with
	   peer VPN gateway. The value should be a CIDR formatted string,
	   for example `192.168.0.0/16`. The ranges should be disjoint.
	   Only IPv4 is supported.
	*/
	LocalTrafficSelectors []string `json:"localTrafficSelectors,omitempty" yaml:"localTrafficSelectors,omitempty"`

	// The interface ID of the external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGatewayInterface int `json:"peerExternalGatewayInterface,omitempty" yaml:"peerExternalGatewayInterface,omitempty"`

	/*
	   URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected.
	   If provided, the VPN tunnel will automatically use the same vpn_gateway_interface
	   ID in the peer GCP VPN gateway.
	   This field must reference a `gcp.compute.HaVpnGateway` resource.
	*/
	PeerGcpGateway string `json:"peerGcpGateway,omitempty" yaml:"peerGcpGateway,omitempty"`

	// IP address of the peer VPN gateway. Only IPv4 is supported.
	PeerIp string `json:"peerIp,omitempty" yaml:"peerIp,omitempty"`

	/*
	   URL of the VPN gateway with which this VPN tunnel is associated.
	   This must be used if a High Availability VPN gateway resource is created.
	   This field must reference a `gcp.compute.HaVpnGateway` resource.
	*/
	VpnGateway string `json:"vpnGateway,omitempty" yaml:"vpnGateway,omitempty"`

	/*
	   Name of the resource. The name must be 1-63 characters long, and
	   comply with RFC1035. Specifically, the name must be 1-63
	   characters long and match the regular expression
	   `a-z?` which means the first character
	   must be a lowercase letter, and all following characters must
	   be a dash, lowercase letter, or digit,
	   except the last character, which cannot be a dash.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The region where the tunnel is located. If unset, is set to the region of `target_vpn_gateway`.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   Shared secret used to set the secure session between the Cloud VPN
	   gateway and the peer VPN gateway.
	   --Note--: This property is sensitive and will not be displayed in the plan.


	   - - -
	*/
	SharedSecret string `json:"sharedSecret,omitempty" yaml:"sharedSecret,omitempty"`

	/*
	   URL of the Target VPN gateway with which this VPN tunnel is
	   associated.
	*/
	TargetVpnGateway string `json:"targetVpnGateway,omitempty" yaml:"targetVpnGateway,omitempty"`

	/*
	   IKE protocol version to use when establishing the VPN tunnel with
	   peer VPN gateway.
	   Acceptable IKE versions are 1 or 2. Default version is 2.
	*/
	IkeVersion int `json:"ikeVersion,omitempty" yaml:"ikeVersion,omitempty"`

	/*
	   Remote traffic selector to use when establishing the VPN tunnel with
	   peer VPN gateway. The value should be a CIDR formatted string,
	   for example `192.168.0.0/16`. The ranges should be disjoint.
	   Only IPv4 is supported.
	*/
	RemoteTrafficSelectors []string `json:"remoteTrafficSelectors,omitempty" yaml:"remoteTrafficSelectors,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Labels to apply to this VpnTunnel.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// URL of the peer side external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGateway string `json:"peerExternalGateway,omitempty" yaml:"peerExternalGateway,omitempty"`

	// URL of router resource to be used for dynamic routing.
	Router string `json:"router,omitempty" yaml:"router,omitempty"`

	// The interface ID of the VPN gateway with which this VPN tunnel is associated.
	VpnGatewayInterface int `json:"vpnGatewayInterface,omitempty" yaml:"vpnGatewayInterface,omitempty"`
}
