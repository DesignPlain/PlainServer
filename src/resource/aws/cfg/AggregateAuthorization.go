package cfg

type AggregateAuthorization struct {
	// Account ID
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// Region
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
