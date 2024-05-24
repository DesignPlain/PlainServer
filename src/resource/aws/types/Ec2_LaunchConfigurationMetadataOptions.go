package types

type Ec2_LaunchConfigurationMetadataOptions struct {
	// The desired HTTP PUT response hop limit for instance metadata requests.
	HttpPutResponseHopLimit int `json:"httpPutResponseHopLimit,omitempty" yaml:"httpPutResponseHopLimit,omitempty"`

	// If session tokens are required: `optional`, `required`.
	HttpTokens string `json:"httpTokens,omitempty" yaml:"httpTokens,omitempty"`

	// The state of the metadata service: `enabled`, `disabled`.
	HttpEndpoint string `json:"httpEndpoint,omitempty" yaml:"httpEndpoint,omitempty"`
}
