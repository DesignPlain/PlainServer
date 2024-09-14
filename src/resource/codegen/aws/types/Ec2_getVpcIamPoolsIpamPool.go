package types

type Ec2_getVpcIamPoolsIpamPool struct {
	// ID of the scope the pool belongs to.
	IpamScopeId string `json:"ipamScopeId,omitempty" yaml:"ipamScopeId,omitempty"`

	// Locale is the Region where your pool is available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region.
	Locale string `json:"locale,omitempty" yaml:"locale,omitempty"`

	// A default netmask length for allocations added to this pool. If, for example, the CIDR assigned to this pool is `10.0.0.0/8` and you enter 16 here, new allocations will default to `10.0.0.0/16`.
	AllocationDefaultNetmaskLength int `json:"allocationDefaultNetmaskLength,omitempty" yaml:"allocationDefaultNetmaskLength,omitempty"`

	// The maximum netmask length that will be required for CIDR allocations in this pool.
	AllocationMaxNetmaskLength int `json:"allocationMaxNetmaskLength,omitempty" yaml:"allocationMaxNetmaskLength,omitempty"`

	// ARN of the pool
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// Limits which service in AWS that the pool can be used in. `ec2` for example, allows users to use space for Elastic IP addresses and VPCs.
	AwsService string `json:"awsService,omitempty" yaml:"awsService,omitempty"`

	// ID of the IPAM pool.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// The minimum netmask length that will be required for CIDR allocations in this pool.
	AllocationMinNetmaskLength int `json:"allocationMinNetmaskLength,omitempty" yaml:"allocationMinNetmaskLength,omitempty"`

	// Tags that are required to create resources in using this pool.
	AllocationResourceTags map[string]string `json:"allocationResourceTags,omitempty" yaml:"allocationResourceTags,omitempty"`

	//
	IpamScopeType string `json:"ipamScopeType,omitempty" yaml:"ipamScopeType,omitempty"`

	// If enabled, IPAM will continuously look for resources within the CIDR range of this pool and automatically import them as allocations into your IPAM.
	AutoImport bool `json:"autoImport,omitempty" yaml:"autoImport,omitempty"`

	//
	PoolDepth int `json:"poolDepth,omitempty" yaml:"poolDepth,omitempty"`

	// Defines whether or not IPv6 pool space is publicly advertisable over the internet.
	PubliclyAdvertisable bool `json:"publiclyAdvertisable,omitempty" yaml:"publiclyAdvertisable,omitempty"`

	// ID of the source IPAM pool.
	SourceIpamPoolId string `json:"sourceIpamPoolId,omitempty" yaml:"sourceIpamPoolId,omitempty"`

	// Map of tags to assigned to the resource.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// IP protocol assigned to this pool.
	AddressFamily string `json:"addressFamily,omitempty" yaml:"addressFamily,omitempty"`

	// Description for the IPAM pool.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	//
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
