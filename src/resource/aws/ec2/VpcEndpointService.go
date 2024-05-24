package ec2

type VpcEndpointService struct {
	// Amazon Resource Names (ARNs) of one or more Gateway Load Balancers for the endpoint service.
	GatewayLoadBalancerArns []string `json:"gatewayLoadBalancerArns,omitempty" yaml:"gatewayLoadBalancerArns,omitempty"`

	// Amazon Resource Names (ARNs) of one or more Network Load Balancers for the endpoint service.
	NetworkLoadBalancerArns []string `json:"networkLoadBalancerArns,omitempty" yaml:"networkLoadBalancerArns,omitempty"`

	// The private DNS name for the service.
	PrivateDnsName string `json:"privateDnsName,omitempty" yaml:"privateDnsName,omitempty"`

	// The supported IP address types. The possible values are `ipv4` and `ipv6`.
	SupportedIpAddressTypes []string `json:"supportedIpAddressTypes,omitempty" yaml:"supportedIpAddressTypes,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - `true` or `false`.
	AcceptanceRequired bool `json:"acceptanceRequired,omitempty" yaml:"acceptanceRequired,omitempty"`

	// The ARNs of one or more principals allowed to discover the endpoint service.
	AllowedPrincipals []string `json:"allowedPrincipals,omitempty" yaml:"allowedPrincipals,omitempty"`
}
