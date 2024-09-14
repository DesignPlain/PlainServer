package route53

import types "libds/aws/types"

type ResolverEndpoint struct {
	// The protocols you want to use for the Route 53 Resolver endpoint. Valid values: `DoH`, `Do53`, `DoH-FIPS`.
	Protocols []string `json:"protocols,omitempty" yaml:"protocols,omitempty"`

	// The Route 53 Resolver endpoint IP address type. Valid values: `IPV4`, `IPV6`, `DUALSTACK`.
	ResolverEndpointType string `json:"resolverEndpointType,omitempty" yaml:"resolverEndpointType,omitempty"`

	// The ID of one or more security groups that you want to use to control access to this VPC.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   The direction of DNS queries to or from the Route 53 Resolver endpoint.
	   Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
	   or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
	*/
	Direction string `json:"direction,omitempty" yaml:"direction,omitempty"`

	/*
	   The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
	   to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
	*/
	IpAddresses []types.Route53_ResolverEndpointIpAddress `json:"ipAddresses,omitempty" yaml:"ipAddresses,omitempty"`

	// The friendly name of the Route 53 Resolver endpoint.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
