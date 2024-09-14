package types

type Ec2_getVpcEndpointDnsOption struct {
	// The DNS records created for the endpoint.
	DnsRecordIpType string `json:"dnsRecordIpType,omitempty" yaml:"dnsRecordIpType,omitempty"`

	// Indicates whether to enable private DNS only for inbound endpoints.
	PrivateDnsOnlyForInboundResolverEndpoint bool `json:"privateDnsOnlyForInboundResolverEndpoint,omitempty" yaml:"privateDnsOnlyForInboundResolverEndpoint,omitempty"`
}
