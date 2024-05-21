package types

type Filestore_getInstanceNetwork struct {
	// A list of IPv4 or IPv6 addresses.
	IpAddresses []string `json:"ipAddresses,omitempty" yaml:"ipAddresses,omitempty"`

	/*
	   IP versions for which the instance has
	   IP addresses assigned. Possible values: ["ADDRESS_MODE_UNSPECIFIED", "MODE_IPV4", "MODE_IPV6"]
	*/
	Modes []string `json:"modes,omitempty" yaml:"modes,omitempty"`

	/*
	   The name of the GCE VPC network to which the
	   instance is connected.
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   A /29 CIDR block that identifies the range of IP
	   addresses reserved for this instance.
	*/
	ReservedIpRange string `json:"reservedIpRange,omitempty" yaml:"reservedIpRange,omitempty"`

	/*
	   The network connect mode of the Filestore instance.
	   If not provided, the connect mode defaults to
	   DIRECT_PEERING. Default value: "DIRECT_PEERING" Possible values: ["DIRECT_PEERING", "PRIVATE_SERVICE_ACCESS"]
	*/
	ConnectMode string `json:"connectMode,omitempty" yaml:"connectMode,omitempty"`
}
