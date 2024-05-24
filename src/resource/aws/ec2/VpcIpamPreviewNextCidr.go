package ec2

type VpcIpamPreviewNextCidr struct {
	// The netmask length of the CIDR you would like to preview from the IPAM pool.
	NetmaskLength int `json:"netmaskLength,omitempty" yaml:"netmaskLength,omitempty"`

	// Exclude a particular CIDR range from being returned by the pool.
	DisallowedCidrs []string `json:"disallowedCidrs,omitempty" yaml:"disallowedCidrs,omitempty"`

	// The ID of the pool to which you want to assign a CIDR.
	IpamPoolId string `json:"ipamPoolId,omitempty" yaml:"ipamPoolId,omitempty"`
}
