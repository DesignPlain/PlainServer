package types

type Vmwareengine_SubnetDhcpAddressRange struct {
	/*
	   (Output)
	   The first IP address of the range.
	*/
	FirstAddress string `json:"firstAddress,omitempty" yaml:"firstAddress,omitempty"`

	/*
	   (Output)
	   The last IP address of the range.
	*/
	LastAddress string `json:"lastAddress,omitempty" yaml:"lastAddress,omitempty"`
}
