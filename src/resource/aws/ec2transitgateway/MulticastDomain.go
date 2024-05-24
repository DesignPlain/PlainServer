package ec2transitgateway

type MulticastDomain struct {
	// Whether to enable Internet Group Management Protocol (IGMP) version 2 for the EC2 Transit Gateway Multicast Domain. Valid values: `disable`, `enable`. Default value: `disable`.
	Igmpv2Support string `json:"igmpv2Support,omitempty" yaml:"igmpv2Support,omitempty"`

	// Whether to enable support for statically configuring multicast group sources for the EC2 Transit Gateway Multicast Domain. Valid values: `disable`, `enable`. Default value: `disable`.
	StaticSourcesSupport string `json:"staticSourcesSupport,omitempty" yaml:"staticSourcesSupport,omitempty"`

	// Key-value tags for the EC2 Transit Gateway Multicast Domain. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// EC2 Transit Gateway identifier. The EC2 Transit Gateway must have `multicast_support` enabled.
	TransitGatewayId string `json:"transitGatewayId,omitempty" yaml:"transitGatewayId,omitempty"`

	// Whether to automatically accept cross-account subnet associations that are associated with the EC2 Transit Gateway Multicast Domain. Valid values: `disable`, `enable`. Default value: `disable`.
	AutoAcceptSharedAssociations string `json:"autoAcceptSharedAssociations,omitempty" yaml:"autoAcceptSharedAssociations,omitempty"`
}
