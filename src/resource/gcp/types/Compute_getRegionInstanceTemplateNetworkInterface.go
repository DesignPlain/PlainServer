package types

type Compute_getRegionInstanceTemplateNetworkInterface struct {
	// An array of IPv6 access configurations for this interface. Currently, only one IPv6 access config, DIRECT_IPV6, is supported. If there is no ipv6AccessConfig specified, then this instance will have no external IPv6 Internet access.
	Ipv6AccessConfigs []Compute_getRegionInstanceTemplateNetworkInterfaceIpv6AccessConfig `json:"ipv6AccessConfigs,omitempty" yaml:"ipv6AccessConfigs,omitempty"`

	// One of EXTERNAL, INTERNAL to indicate whether the IP can be accessed from the Internet. This field is always inherited from its subnetwork.
	Ipv6AccessType string `json:"ipv6AccessType,omitempty" yaml:"ipv6AccessType,omitempty"`

	/*
	   The name or self_link of the network to attach this interface to.
	   Use `network` attribute for Legacy or Auto subnetted networks and
	   `subnetwork` for custom subnetted networks.
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	// The type of vNIC to be used on this interface. Possible values:GVNIC, VIRTIO_NET
	NicType string `json:"nicType,omitempty" yaml:"nicType,omitempty"`

	/*
	   The private IP address to assign to the instance. If
	   empty, the address will be automatically assigned.
	*/
	NetworkIp string `json:"networkIp,omitempty" yaml:"networkIp,omitempty"`

	// The networking queue count that's specified by users for the network interface. Both Rx and Tx queues will be set to this number. It will be empty if not specified.
	QueueCount int `json:"queueCount,omitempty" yaml:"queueCount,omitempty"`

	// The stack type for this network interface to identify whether the IPv6 feature is enabled or not. If not specified, IPV4_ONLY will be used.
	StackType string `json:"stackType,omitempty" yaml:"stackType,omitempty"`

	//
	AccessConfigs []Compute_getRegionInstanceTemplateNetworkInterfaceAccessConfig `json:"accessConfigs,omitempty" yaml:"accessConfigs,omitempty"`

	/*
	   An
	   array of alias IP ranges for this network interface. Can only be specified for network
	   interfaces on subnet-mode networks. Structure documented below.
	*/
	AliasIpRanges []Compute_getRegionInstanceTemplateNetworkInterfaceAliasIpRange `json:"aliasIpRanges,omitempty" yaml:"aliasIpRanges,omitempty"`

	// The prefix length of the primary internal IPv6 range.
	InternalIpv6PrefixLength int `json:"internalIpv6PrefixLength,omitempty" yaml:"internalIpv6PrefixLength,omitempty"`

	// An IPv6 internal network address for this network interface. If not specified, Google Cloud will automatically assign an internal IPv6 address from the instance's subnetwork.
	Ipv6Address string `json:"ipv6Address,omitempty" yaml:"ipv6Address,omitempty"`

	// The name of the instance template. One of `name` or `filter` must be provided.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   the name of the subnetwork to attach this interface
	   to. The subnetwork must exist in the same `region` this instance will be
	   created in. Either `network` or `subnetwork` must be provided.
	*/
	Subnetwork string `json:"subnetwork,omitempty" yaml:"subnetwork,omitempty"`

	/*
	   The ID of the project in which the subnetwork belongs.
	   If it is not provided, the provider project is used.
	*/
	SubnetworkProject string `json:"subnetworkProject,omitempty" yaml:"subnetworkProject,omitempty"`
}
