package types

type S3_BucketLifecycleConfigurationV2RuleFilter struct {
	// Configuration block for specifying a tag key and value. See below.
	Tag S3_BucketLifecycleConfigurationV2RuleFilterTag `json:"tag,omitempty" yaml:"tag,omitempty"`

	// Configuration block used to apply a logical `AND` to two or more predicates. See below. The Lifecycle Rule will apply to any object matching all the predicates configured inside the `and` block.
	And S3_BucketLifecycleConfigurationV2RuleFilterAnd `json:"and,omitempty" yaml:"and,omitempty"`

	// Minimum object size (in bytes) to which the rule applies.
	ObjectSizeGreaterThan string `json:"objectSizeGreaterThan,omitempty" yaml:"objectSizeGreaterThan,omitempty"`

	// Maximum object size (in bytes) to which the rule applies.
	ObjectSizeLessThan string `json:"objectSizeLessThan,omitempty" yaml:"objectSizeLessThan,omitempty"`

	// Prefix identifying one or more objects to which the rule applies. Defaults to an empty string (`""`) if not specified.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
}
