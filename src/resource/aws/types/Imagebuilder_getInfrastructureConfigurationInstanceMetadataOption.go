package types

type Imagebuilder_getInfrastructureConfigurationInstanceMetadataOption struct {
	// Number of hops that an instance can traverse to reach its destonation.
	HttpPutResponseHopLimit int `json:"httpPutResponseHopLimit,omitempty" yaml:"httpPutResponseHopLimit,omitempty"`

	// Whether a signed token is required for instance metadata retrieval requests.
	HttpTokens string `json:"httpTokens,omitempty" yaml:"httpTokens,omitempty"`
}
