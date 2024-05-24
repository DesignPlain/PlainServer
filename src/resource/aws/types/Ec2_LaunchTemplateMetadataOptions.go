package types

type Ec2_LaunchTemplateMetadataOptions struct {
	// The desired HTTP PUT response hop limit for instance metadata requests. The larger the number, the further instance metadata requests can travel. Can be an integer from `1` to `64`. (Default: `1`).
	HttpPutResponseHopLimit int `json:"httpPutResponseHopLimit,omitempty" yaml:"httpPutResponseHopLimit,omitempty"`

	// Whether or not the metadata service requires session tokens, also referred to as _Instance Metadata Service Version 2 (IMDSv2)_. Can be `"optional"` or `"required"`. (Default: `"optional"`).
	HttpTokens string `json:"httpTokens,omitempty" yaml:"httpTokens,omitempty"`

	/*
	   Enables or disables access to instance tags from the instance metadata service. Can be `"enabled"` or `"disabled"`.

	   For more information, see the documentation on the [Instance Metadata Service](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html).
	*/
	InstanceMetadataTags string `json:"instanceMetadataTags,omitempty" yaml:"instanceMetadataTags,omitempty"`

	// Whether the metadata service is available. Can be `"enabled"` or `"disabled"`. (Default: `"enabled"`).
	HttpEndpoint string `json:"httpEndpoint,omitempty" yaml:"httpEndpoint,omitempty"`

	// Enables or disables the IPv6 endpoint for the instance metadata service. Can be `"enabled"` or `"disabled"`.
	HttpProtocolIpv6 string `json:"httpProtocolIpv6,omitempty" yaml:"httpProtocolIpv6,omitempty"`
}
