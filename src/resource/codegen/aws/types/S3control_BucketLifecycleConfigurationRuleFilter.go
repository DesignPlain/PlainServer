package types

type S3control_BucketLifecycleConfigurationRuleFilter struct {
	// Object prefix for rule filtering.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// Key-value map of object tags for rule filtering.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
