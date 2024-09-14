package types

type Ec2_getLaunchTemplateMetadataOption struct {
	//
	HttpEndpoint string `json:"httpEndpoint,omitempty" yaml:"httpEndpoint,omitempty"`

	//
	HttpProtocolIpv6 string `json:"httpProtocolIpv6,omitempty" yaml:"httpProtocolIpv6,omitempty"`

	//
	HttpPutResponseHopLimit int `json:"httpPutResponseHopLimit,omitempty" yaml:"httpPutResponseHopLimit,omitempty"`

	//
	HttpTokens string `json:"httpTokens,omitempty" yaml:"httpTokens,omitempty"`

	//
	InstanceMetadataTags string `json:"instanceMetadataTags,omitempty" yaml:"instanceMetadataTags,omitempty"`
}
