package types

type Vmwareengine_getSubnetDhcpAddressRange struct {
	// The first IP address of the range.
	FirstAddress string `json:"firstAddress,omitempty" yaml:"firstAddress,omitempty"`

	// The last IP address of the range.
	LastAddress string `json:"lastAddress,omitempty" yaml:"lastAddress,omitempty"`
}
