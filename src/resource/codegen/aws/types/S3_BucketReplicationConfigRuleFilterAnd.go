package types

type S3_BucketReplicationConfigRuleFilterAnd struct {
	// Object key name prefix that identifies subset of objects to which the rule applies. Must be less than or equal to 1024 characters in length.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// Map of tags (key and value pairs) that identifies a subset of objects to which the rule applies. The rule applies only to objects having all the tags in its tagset.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
