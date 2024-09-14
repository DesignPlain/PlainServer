package ec2

type InstanceMetadataDefaults struct {
	// Whether the metadata service is available. Can be `"enabled"`, `"disabled"`, or `"no-preference"`. Default: `"no-preference"`.
	HttpEndpoint string `json:"httpEndpoint,omitempty" yaml:"httpEndpoint,omitempty"`

	// The desired HTTP PUT response hop limit for instance metadata requests. The larger the number, the further instance metadata requests can travel. Can be an integer from `1` to `64`, or `-1` to indicate no preference. Default: `-1`.
	HttpPutResponseHopLimit int `json:"httpPutResponseHopLimit,omitempty" yaml:"httpPutResponseHopLimit,omitempty"`

	// Whether the metadata service requires session tokens, also referred to as _Instance Metadata Service Version 2 (IMDSv2)_. Can be `"optional"`, `"required"`, or `"no-preference"`. Default: `"no-preference"`.
	HttpTokens string `json:"httpTokens,omitempty" yaml:"httpTokens,omitempty"`

	// Enables or disables access to instance tags from the instance metadata service. Can be `"enabled"`, `"disabled"`, or `"no-preference"`. Default: `"no-preference"`.
	InstanceMetadataTags string `json:"instanceMetadataTags,omitempty" yaml:"instanceMetadataTags,omitempty"`
}
