package types

type S3_BucketLifecycleConfigurationV2RuleFilterAnd struct {
	// Minimum object size to which the rule applies. Value must be at least `0` if specified.
	ObjectSizeGreaterThan int `json:"objectSizeGreaterThan,omitempty" yaml:"objectSizeGreaterThan,omitempty"`

	// Maximum object size to which the rule applies. Value must be at least `1` if specified.
	ObjectSizeLessThan int `json:"objectSizeLessThan,omitempty" yaml:"objectSizeLessThan,omitempty"`

	// Prefix identifying one or more objects to which the rule applies.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// Key-value map of resource tags. All of these tags must exist in the object's tag set in order for the rule to apply.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
