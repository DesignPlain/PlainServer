package types

type Imagebuilder_InfrastructureConfigurationInstanceMetadataOptions struct {
	// Whether a signed token is required for instance metadata retrieval requests. Valid values: `required`, `optional`.
	HttpTokens string `json:"httpTokens,omitempty" yaml:"httpTokens,omitempty"`

	// The number of hops that an instance can traverse to reach its destonation.
	HttpPutResponseHopLimit int `json:"httpPutResponseHopLimit,omitempty" yaml:"httpPutResponseHopLimit,omitempty"`
}
