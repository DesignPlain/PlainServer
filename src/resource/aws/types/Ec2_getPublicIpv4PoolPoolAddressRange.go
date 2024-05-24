package types

type Ec2_getPublicIpv4PoolPoolAddressRange struct {
	// Number of addresses in the range.
	AddressCount int `json:"addressCount,omitempty" yaml:"addressCount,omitempty"`

	// Number of available addresses in the range.
	AvailableAddressCount int `json:"availableAddressCount,omitempty" yaml:"availableAddressCount,omitempty"`

	// First address in the range.
	FirstAddress string `json:"firstAddress,omitempty" yaml:"firstAddress,omitempty"`

	// Last address in the range.
	LastAddress string `json:"lastAddress,omitempty" yaml:"lastAddress,omitempty"`
}
