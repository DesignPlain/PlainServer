package ec2

type DefaultVpc struct {
	//
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	EnableDnsSupport bool `json:"enableDnsSupport,omitempty" yaml:"enableDnsSupport,omitempty"`

	//
	EnableNetworkAddressUsageMetrics bool `json:"enableNetworkAddressUsageMetrics,omitempty" yaml:"enableNetworkAddressUsageMetrics,omitempty"`

	// Whether destroying the resource deletes the default VPC. Default: `false`
	ForceDestroy bool `json:"forceDestroy,omitempty" yaml:"forceDestroy,omitempty"`

	//
	Ipv6CidrBlock string `json:"ipv6CidrBlock,omitempty" yaml:"ipv6CidrBlock,omitempty"`

	//
	Ipv6CidrBlockNetworkBorderGroup string `json:"ipv6CidrBlockNetworkBorderGroup,omitempty" yaml:"ipv6CidrBlockNetworkBorderGroup,omitempty"`

	//
	Ipv6NetmaskLength int `json:"ipv6NetmaskLength,omitempty" yaml:"ipv6NetmaskLength,omitempty"`

	//
	AssignGeneratedIpv6CidrBlock bool `json:"assignGeneratedIpv6CidrBlock,omitempty" yaml:"assignGeneratedIpv6CidrBlock,omitempty"`

	//
	EnableDnsHostnames bool `json:"enableDnsHostnames,omitempty" yaml:"enableDnsHostnames,omitempty"`

	//
	Ipv6IpamPoolId string `json:"ipv6IpamPoolId,omitempty" yaml:"ipv6IpamPoolId,omitempty"`
}
