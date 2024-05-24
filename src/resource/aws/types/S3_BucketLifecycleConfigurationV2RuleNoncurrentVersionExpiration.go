package types

type S3_BucketLifecycleConfigurationV2RuleNoncurrentVersionExpiration struct {
	// Number of noncurrent versions Amazon S3 will retain. Must be a non-zero positive integer.
	NewerNoncurrentVersions string `json:"newerNoncurrentVersions,omitempty" yaml:"newerNoncurrentVersions,omitempty"`

	// Number of days an object is noncurrent before Amazon S3 can perform the associated action. Must be a positive integer.
	NoncurrentDays int `json:"noncurrentDays,omitempty" yaml:"noncurrentDays,omitempty"`
}
