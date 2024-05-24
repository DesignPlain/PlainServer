package types

type Ec2_VpcEndpointDnsOptions struct {
	// The DNS records created for the endpoint. Valid values are `ipv4`, `dualstack`, `service-defined`, and `ipv6`.
	DnsRecordIpType string `json:"dnsRecordIpType,omitempty" yaml:"dnsRecordIpType,omitempty"`

	// Indicates whether to enable private DNS only for inbound endpoints. This option is available only for services that support both gateway and interface endpoints. It routes traffic that originates from the VPC to the gateway endpoint and traffic that originates from on-premises to the interface endpoint. Default is `false`. Can only be specified if private_dns_enabled is `true`.
	PrivateDnsOnlyForInboundResolverEndpoint bool `json:"privateDnsOnlyForInboundResolverEndpoint,omitempty" yaml:"privateDnsOnlyForInboundResolverEndpoint,omitempty"`
}
