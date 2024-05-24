package types

type S3_BucketReplicationConfigRuleFilter struct {
	// Configuration block for specifying a tag key and value. See below.
	Tag S3_BucketReplicationConfigRuleFilterTag `json:"tag,omitempty" yaml:"tag,omitempty"`

	// Configuration block for specifying rule filters. This element is required only if you specify more than one filter. See and below for more details.
	And S3_BucketReplicationConfigRuleFilterAnd `json:"and,omitempty" yaml:"and,omitempty"`

	// Object key name prefix that identifies subset of objects to which the rule applies. Must be less than or equal to 1024 characters in length.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
}
