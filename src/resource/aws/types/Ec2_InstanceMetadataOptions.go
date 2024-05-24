package types

type Ec2_InstanceMetadataOptions struct {
	// Whether or not the metadata service requires session tokens, also referred to as _Instance Metadata Service Version 2 (IMDSv2)_. Valid values include `optional` or `required`. Defaults to `optional`.
	HttpTokens string `json:"httpTokens,omitempty" yaml:"httpTokens,omitempty"`

	/*
	   Enables or disables access to instance tags from the instance metadata service. Valid values include `enabled` or `disabled`. Defaults to `disabled`.

	   For more information, see the documentation on the [Instance Metadata Service](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html).
	*/
	InstanceMetadataTags string `json:"instanceMetadataTags,omitempty" yaml:"instanceMetadataTags,omitempty"`

	// Whether the metadata service is available. Valid values include `enabled` or `disabled`. Defaults to `enabled`.
	HttpEndpoint string `json:"httpEndpoint,omitempty" yaml:"httpEndpoint,omitempty"`

	// Whether the IPv6 endpoint for the instance metadata service is enabled. Defaults to `disabled`.
	HttpProtocolIpv6 string `json:"httpProtocolIpv6,omitempty" yaml:"httpProtocolIpv6,omitempty"`

	// Desired HTTP PUT response hop limit for instance metadata requests. The larger the number, the further instance metadata requests can travel. Valid values are integer from `1` to `64`. Defaults to `1`.
	HttpPutResponseHopLimit int `json:"httpPutResponseHopLimit,omitempty" yaml:"httpPutResponseHopLimit,omitempty"`
}
