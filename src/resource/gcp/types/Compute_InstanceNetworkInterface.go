package types

type Compute_InstanceNetworkInterface struct {
	/*
	   One of EXTERNAL, INTERNAL to indicate whether the IP can be accessed from the Internet.
	   This field is always inherited from its subnetwork.
	*/
	Ipv6AccessType string `json:"ipv6AccessType,omitempty" yaml:"ipv6AccessType,omitempty"`

	// The URL of the network attachment that this interface should connect to in the following format: `projects/{projectNumber}/regions/{region_name}/networkAttachments/{network_attachment_name}`.
	NetworkAttachment string `json:"networkAttachment,omitempty" yaml:"networkAttachment,omitempty"`

	/*
	   The name or self_link of the subnetwork to attach this
	   interface to. Either `network` or `subnetwork` must be provided. If network isn't provided
	   it will be inferred from the subnetwork. The subnetwork must exist in the same region this
	   instance will be created in. If the network resource is in
	   [legacy](https://cloud.google.com/vpc/docs/legacy) mode, do not specify this field. If the
	   network is in auto subnet mode, specifying the subnetwork is optional. If the network is
	   in custom subnet mode, specifying the subnetwork is required.
	*/
	Subnetwork string `json:"subnetwork,omitempty" yaml:"subnetwork,omitempty"`

	/*
	   An
	   array of alias IP ranges for this network interface. Can only be specified for network
	   interfaces on subnet-mode networks. Structure documented below.
	*/
	AliasIpRanges []Compute_InstanceNetworkInterfaceAliasIpRange `json:"aliasIpRanges,omitempty" yaml:"aliasIpRanges,omitempty"`

	// The prefix length of the primary internal IPv6 range.
	InternalIpv6PrefixLength int `json:"internalIpv6PrefixLength,omitempty" yaml:"internalIpv6PrefixLength,omitempty"`

	// The stack type for this network interface to identify whether the IPv6 feature is enabled or not. Values are IPV4_IPV6 or IPV4_ONLY. If not specified, IPV4_ONLY will be used.
	StackType string `json:"stackType,omitempty" yaml:"stackType,omitempty"`

	/*
	   The project in which the subnetwork belongs.
	   If the `subnetwork` is a self_link, this field is ignored in favor of the project
	   defined in the subnetwork self_link. If the `subnetwork` is a name and this
	   field is not provided, the provider project is used.
	*/
	SubnetworkProject string `json:"subnetworkProject,omitempty" yaml:"subnetworkProject,omitempty"`

	/*
	   The name or self_link of the network to attach this interface to.
	   Either `network` or `subnetwork` must be provided. If network isn't provided it will
	   be inferred from the subnetwork.
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   The private IP address to assign to the instance. If
	   empty, the address will be automatically assigned.
	*/
	NetworkIp string `json:"networkIp,omitempty" yaml:"networkIp,omitempty"`

	// A full or partial URL to a security policy to add to this instance. If this field is set to an empty string it will remove the associated security policy.
	SecurityPolicy string `json:"securityPolicy,omitempty" yaml:"securityPolicy,omitempty"`

	// An IPv6 internal network address for this network interface. If not specified, Google Cloud will automatically assign an internal IPv6 address from the instance's subnetwork.
	Ipv6Address string `json:"ipv6Address,omitempty" yaml:"ipv6Address,omitempty"`

	/*
	   A unique name for the resource, required by GCE.
	   Changing this forces a new resource to be created.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The type of vNIC to be used on this interface. Possible values: GVNIC, VIRTIO_NET.
	NicType string `json:"nicType,omitempty" yaml:"nicType,omitempty"`

	// The networking queue count that's specified by users for the network interface. Both Rx and Tx queues will be set to this number. It will be empty if not specified.
	QueueCount int `json:"queueCount,omitempty" yaml:"queueCount,omitempty"`

	// Access configurations, i.e. IPs via which this instance can be accessed via the Internet.
	AccessConfigs []Compute_InstanceNetworkInterfaceAccessConfig `json:"accessConfigs,omitempty" yaml:"accessConfigs,omitempty"`

	/*
	   An array of IPv6 access configurations for this interface.
	   Currently, only one IPv6 access config, DIRECT_IPV6, is supported. If there is no ipv6AccessConfig
	   specified, then this instance will have no external IPv6 Internet access. Structure documented below.
	*/
	Ipv6AccessConfigs []Compute_InstanceNetworkInterfaceIpv6AccessConfig `json:"ipv6AccessConfigs,omitempty" yaml:"ipv6AccessConfigs,omitempty"`
}
