package compute

type RouterInterface struct {
	/*
	   The ID of the project in which this interface's routerbelongs.
	   If it is not provided, the provider project is used. Changing this forces a new interface to be created.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The name of the interface that is redundant to
	   this interface. Changing this forces a new interface to be created.
	*/
	RedundantInterface string `json:"redundantInterface,omitempty" yaml:"redundantInterface,omitempty"`

	/*
	   The region this interface's router sits in.
	   If not specified, the project region will be used. Changing this forces a new interface to be created.
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   The name of the router this interface will be attached to.
	   Changing this forces a new interface to be created.

	   In addition to the above required fields, a router interface must have specified either `ip_range` or exactly one of `vpn_tunnel`, `interconnect_attachment` or `subnetwork`, or both.

	   - - -
	*/
	Router string `json:"router,omitempty" yaml:"router,omitempty"`

	/*
	   The URI of the subnetwork resource that this interface
	   belongs to, which must be in the same region as the Cloud Router. When you establish a BGP session to a VM instance using this interface, the VM instance must belong to the same subnetwork as the subnetwork specified here. Changing this forces a new interface to be created. Only one of `vpn_tunnel`, `interconnect_attachment` or `subnetwork` can be specified.
	*/
	Subnetwork string `json:"subnetwork,omitempty" yaml:"subnetwork,omitempty"`

	/*
	   The name or resource link to the VPN tunnel this
	   interface will be linked to. Changing this forces a new interface to be created. Only
	   one of `vpn_tunnel`, `interconnect_attachment` or `subnetwork` can be specified.
	*/
	VpnTunnel string `json:"vpnTunnel,omitempty" yaml:"vpnTunnel,omitempty"`

	/*
	   The name or resource link to the
	   VLAN interconnect for this interface. Changing this forces a new interface to
	   be created. Only one of `vpn_tunnel`, `interconnect_attachment` or `subnetwork` can be specified.
	*/
	InterconnectAttachment string `json:"interconnectAttachment,omitempty" yaml:"interconnectAttachment,omitempty"`

	/*
	   A unique name for the interface, required by GCE. Changing
	   this forces a new interface to be created.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   IP address and range of the interface. The IP range must be
	   in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
	*/
	IpRange string `json:"ipRange,omitempty" yaml:"ipRange,omitempty"`

	/*
	   The regional private internal IP address that is used
	   to establish BGP sessions to a VM instance acting as a third-party Router Appliance. Changing this forces a new interface to be created.
	*/
	PrivateIpAddress string `json:"privateIpAddress,omitempty" yaml:"privateIpAddress,omitempty"`
}
