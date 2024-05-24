package ec2

type Subnet struct {
	// AZ ID of the subnet. This argument is not supported in all regions or partitions. If necessary, use `availability_zone` instead.
	AvailabilityZoneId string `json:"availabilityZoneId,omitempty" yaml:"availabilityZoneId,omitempty"`

	// The customer owned IPv4 address pool. Typically used with the `map_customer_owned_ip_on_launch` argument. The `outpost_arn` argument must be specified when configured.
	CustomerOwnedIpv4Pool string `json:"customerOwnedIpv4Pool,omitempty" yaml:"customerOwnedIpv4Pool,omitempty"`

	// Indicates the device position for local network interfaces in this subnet. For example, 1 indicates local network interfaces in this subnet are the secondary network interface (eth1). A local network interface cannot be the primary network interface (eth0).
	EnableLniAtDeviceIndex int `json:"enableLniAtDeviceIndex,omitempty" yaml:"enableLniAtDeviceIndex,omitempty"`

	// Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records. Default: `false`.
	EnableResourceNameDnsAaaaRecordOnLaunch bool `json:"enableResourceNameDnsAaaaRecordOnLaunch,omitempty" yaml:"enableResourceNameDnsAaaaRecordOnLaunch,omitempty"`

	// The type of hostnames to assign to instances in the subnet at launch. For IPv6-only subnets, an instance DNS name must be based on the instance ID. For dual-stack and IPv4-only subnets, you can specify whether DNS names use the instance IPv4 address or the instance ID. Valid values: `ip-name`, `resource-name`.
	PrivateDnsHostnameTypeOnLaunch string `json:"privateDnsHostnameTypeOnLaunch,omitempty" yaml:"privateDnsHostnameTypeOnLaunch,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specify `true` to indicate that network interfaces created in the subnet should be assigned a customer owned IP address. The `customer_owned_ipv4_pool` and `outpost_arn` arguments must be specified when set to `true`. Default is `false`.
	MapCustomerOwnedIpOnLaunch bool `json:"mapCustomerOwnedIpOnLaunch,omitempty" yaml:"mapCustomerOwnedIpOnLaunch,omitempty"`

	// The VPC ID.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	/*
	   Specify true to indicate
	   that network interfaces created in the specified subnet should be
	   assigned an IPv6 address. Default is `false`
	*/
	AssignIpv6AddressOnCreation bool `json:"assignIpv6AddressOnCreation,omitempty" yaml:"assignIpv6AddressOnCreation,omitempty"`

	// AZ for the subnet.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations. Default: `false`.
	EnableDns64 bool `json:"enableDns64,omitempty" yaml:"enableDns64,omitempty"`

	// Indicates whether to respond to DNS queries for instance hostnames with DNS A records. Default: `false`.
	EnableResourceNameDnsARecordOnLaunch bool `json:"enableResourceNameDnsARecordOnLaunch,omitempty" yaml:"enableResourceNameDnsARecordOnLaunch,omitempty"`

	/*
	   The IPv6 network range for the subnet,
	   in CIDR notation. The subnet size must use a /64 prefix length.
	*/
	Ipv6CidrBlock string `json:"ipv6CidrBlock,omitempty" yaml:"ipv6CidrBlock,omitempty"`

	// Indicates whether to create an IPv6-only subnet. Default: `false`.
	Ipv6Native bool `json:"ipv6Native,omitempty" yaml:"ipv6Native,omitempty"`

	// The IPv4 CIDR block for the subnet.
	CidrBlock string `json:"cidrBlock,omitempty" yaml:"cidrBlock,omitempty"`

	/*
	   Specify true to indicate
	   that instances launched into the subnet should be assigned
	   a public IP address. Default is `false`.
	*/
	MapPublicIpOnLaunch bool `json:"mapPublicIpOnLaunch,omitempty" yaml:"mapPublicIpOnLaunch,omitempty"`

	// The Amazon Resource Name (ARN) of the Outpost.
	OutpostArn string `json:"outpostArn,omitempty" yaml:"outpostArn,omitempty"`
}
