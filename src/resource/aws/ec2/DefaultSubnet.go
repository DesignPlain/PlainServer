package ec2

type DefaultSubnet struct {
	//
	AssignIpv6AddressOnCreation bool `json:"assignIpv6AddressOnCreation,omitempty" yaml:"assignIpv6AddressOnCreation,omitempty"`

	/*
	   is required
	   - The `availability_zone_id`, `cidr_block` and `vpc_id` arguments become computed attributes
	   - The default value for `map_public_ip_on_launch` is `true`

	   This resource supports the following additional arguments:
	*/
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	//
	EnableResourceNameDnsARecordOnLaunch bool `json:"enableResourceNameDnsARecordOnLaunch,omitempty" yaml:"enableResourceNameDnsARecordOnLaunch,omitempty"`

	//
	EnableResourceNameDnsAaaaRecordOnLaunch bool `json:"enableResourceNameDnsAaaaRecordOnLaunch,omitempty" yaml:"enableResourceNameDnsAaaaRecordOnLaunch,omitempty"`

	//
	Ipv6Native bool `json:"ipv6Native,omitempty" yaml:"ipv6Native,omitempty"`

	//
	MapCustomerOwnedIpOnLaunch bool `json:"mapCustomerOwnedIpOnLaunch,omitempty" yaml:"mapCustomerOwnedIpOnLaunch,omitempty"`

	//
	MapPublicIpOnLaunch bool `json:"mapPublicIpOnLaunch,omitempty" yaml:"mapPublicIpOnLaunch,omitempty"`

	//
	PrivateDnsHostnameTypeOnLaunch string `json:"privateDnsHostnameTypeOnLaunch,omitempty" yaml:"privateDnsHostnameTypeOnLaunch,omitempty"`

	//
	CustomerOwnedIpv4Pool string `json:"customerOwnedIpv4Pool,omitempty" yaml:"customerOwnedIpv4Pool,omitempty"`

	//
	EnableDns64 bool `json:"enableDns64,omitempty" yaml:"enableDns64,omitempty"`

	// Whether destroying the resource deletes the default subnet. Default: `false`
	ForceDestroy bool `json:"forceDestroy,omitempty" yaml:"forceDestroy,omitempty"`

	//
	Ipv6CidrBlock string `json:"ipv6CidrBlock,omitempty" yaml:"ipv6CidrBlock,omitempty"`

	//
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
