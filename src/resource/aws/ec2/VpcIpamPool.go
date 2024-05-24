package ec2

type VpcIpamPool struct {
	// A default netmask length for allocations added to this pool. If, for example, the CIDR assigned to this pool is 10.0.0.0/8 and you enter 16 here, new allocations will default to 10.0.0.0/16 (unless you provide a different netmask value when you create the new allocation).
	AllocationDefaultNetmaskLength int `json:"allocationDefaultNetmaskLength,omitempty" yaml:"allocationDefaultNetmaskLength,omitempty"`

	// Tags that are required for resources that use CIDRs from this IPAM pool. Resources that do not have these tags will not be allowed to allocate space from the pool. If the resources have their tags changed after they have allocated space or if the allocation tagging requirements are changed on the pool, the resource may be marked as noncompliant.
	AllocationResourceTags map[string]string `json:"allocationResourceTags,omitempty" yaml:"allocationResourceTags,omitempty"`

	/*
	   If you include this argument, IPAM automatically imports any VPCs you have in your scope that fall
	   within the CIDR range in the pool.
	*/
	AutoImport bool `json:"autoImport,omitempty" yaml:"autoImport,omitempty"`

	// A description for the IPAM pool.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The ID of the scope in which you would like to create the IPAM pool.
	IpamScopeId string `json:"ipamScopeId,omitempty" yaml:"ipamScopeId,omitempty"`

	// The IP protocol assigned to this pool. You must choose either IPv4 or IPv6 protocol for a pool.
	AddressFamily string `json:"addressFamily,omitempty" yaml:"addressFamily,omitempty"`

	// The maximum netmask length that will be required for CIDR allocations in this pool.
	AllocationMaxNetmaskLength int `json:"allocationMaxNetmaskLength,omitempty" yaml:"allocationMaxNetmaskLength,omitempty"`

	// Defines whether or not IPv6 pool space is publicly advertisable over the internet. This argument is required if `address_family = "ipv6"` and `public_ip_source = "byoip"`, default is `false`. This option is not available for IPv4 pool space or if `public_ip_source = "amazon"`.
	PubliclyAdvertisable bool `json:"publiclyAdvertisable,omitempty" yaml:"publiclyAdvertisable,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The minimum netmask length that will be required for CIDR allocations in this pool.
	AllocationMinNetmaskLength int `json:"allocationMinNetmaskLength,omitempty" yaml:"allocationMinNetmaskLength,omitempty"`

	// The locale in which you would like to create the IPAM pool. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. Possible values: Any AWS region, such as `us-east-1`.
	Locale string `json:"locale,omitempty" yaml:"locale,omitempty"`

	// The ID of the source IPAM pool. Use this argument to create a child pool within an existing pool.
	SourceIpamPoolId string `json:"sourceIpamPoolId,omitempty" yaml:"sourceIpamPoolId,omitempty"`

	// Limits which AWS service the pool can be used in. Only useable on public scopes. Valid Values: `ec2`.
	AwsService string `json:"awsService,omitempty" yaml:"awsService,omitempty"`

	// The IP address source for pools in the public scope. Only used for provisioning IP address CIDRs to pools in the public scope. Valid values are `byoip` or `amazon`. Default is `byoip`.
	PublicIpSource string `json:"publicIpSource,omitempty" yaml:"publicIpSource,omitempty"`
}
