package types

type Ec2_getLaunchConfigurationMetadataOption struct {
	// State of the metadata service: `enabled`, `disabled`.
	HttpEndpoint string `json:"httpEndpoint,omitempty" yaml:"httpEndpoint,omitempty"`

	// The desired HTTP PUT response hop limit for instance metadata requests.
	HttpPutResponseHopLimit int `json:"httpPutResponseHopLimit,omitempty" yaml:"httpPutResponseHopLimit,omitempty"`

	// If session tokens are required: `optional`, `required`.
	HttpTokens string `json:"httpTokens,omitempty" yaml:"httpTokens,omitempty"`
}
