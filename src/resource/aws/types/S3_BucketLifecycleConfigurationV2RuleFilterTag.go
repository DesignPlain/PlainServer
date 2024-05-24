package types

type S3_BucketLifecycleConfigurationV2RuleFilterTag struct {
	// Name of the object key.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Value of the tag.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
