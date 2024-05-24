package types

type Ec2_getInstanceMetadataOption struct {
	// State of the metadata service: `enabled`, `disabled`.
	HttpEndpoint string `json:"httpEndpoint,omitempty" yaml:"httpEndpoint,omitempty"`

	// Whether the IPv6 endpoint for the instance metadata service is `enabled` or `disabled`
	HttpProtocolIpv6 string `json:"httpProtocolIpv6,omitempty" yaml:"httpProtocolIpv6,omitempty"`

	// Desired HTTP PUT response hop limit for instance metadata requests.
	HttpPutResponseHopLimit int `json:"httpPutResponseHopLimit,omitempty" yaml:"httpPutResponseHopLimit,omitempty"`

	// If session tokens are required: `optional`, `required`.
	HttpTokens string `json:"httpTokens,omitempty" yaml:"httpTokens,omitempty"`

	// If access to instance tags is allowed from the metadata service: `enabled`, `disabled`.
	InstanceMetadataTags string `json:"instanceMetadataTags,omitempty" yaml:"instanceMetadataTags,omitempty"`
}
