package types

type S3_BucketIntelligentTieringConfigurationFilter struct {
	// Object key name prefix that identifies the subset of objects to which the configuration applies.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// All of these tags must exist in the object's tag set in order for the configuration to apply.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
