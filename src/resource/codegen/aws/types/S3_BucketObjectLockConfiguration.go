package types

type S3_BucketObjectLockConfiguration struct {
	// Indicates whether this bucket has an Object Lock configuration enabled. Valid values are `true` or `false`. This argument is not supported in all regions or partitions.
	ObjectLockEnabled string `json:"objectLockEnabled,omitempty" yaml:"objectLockEnabled,omitempty"`

	// Object Lock rule in place for this bucket (documented below).
	Rule S3_BucketObjectLockConfigurationRule `json:"rule,omitempty" yaml:"rule,omitempty"`
}
