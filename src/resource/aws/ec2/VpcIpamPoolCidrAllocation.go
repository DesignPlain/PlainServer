package ec2

type VpcIpamPoolCidrAllocation struct {
	// Exclude a particular CIDR range from being returned by the pool.
	DisallowedCidrs []string `json:"disallowedCidrs,omitempty" yaml:"disallowedCidrs,omitempty"`

	// The ID of the pool to which you want to assign a CIDR.
	IpamPoolId string `json:"ipamPoolId,omitempty" yaml:"ipamPoolId,omitempty"`

	// The netmask length of the CIDR you would like to allocate to the IPAM pool. Valid Values: `0-128`.
	NetmaskLength int `json:"netmaskLength,omitempty" yaml:"netmaskLength,omitempty"`

	// The CIDR you want to assign to the pool.
	Cidr string `json:"cidr,omitempty" yaml:"cidr,omitempty"`

	// The description for the allocation.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
