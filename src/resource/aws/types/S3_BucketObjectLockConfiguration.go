package types

type S3_BucketObjectLockConfiguration struct {
	// Indicates whether this bucket has an Object Lock configuration enabled. Valid value is `Enabled`.
	ObjectLockEnabled string `json:"objectLockEnabled,omitempty" yaml:"objectLockEnabled,omitempty"`

	// The Object Lock rule in place for this bucket.
	Rule S3_BucketObjectLockConfigurationRule `json:"rule,omitempty" yaml:"rule,omitempty"`
}
