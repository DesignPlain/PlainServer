package types

type S3_BucketLifecycleConfigurationV2RuleNoncurrentVersionTransition struct {
	// Number of noncurrent versions Amazon S3 will retain. Must be a non-zero positive integer.
	NewerNoncurrentVersions string `json:"newerNoncurrentVersions,omitempty" yaml:"newerNoncurrentVersions,omitempty"`

	// Number of days an object is noncurrent before Amazon S3 can perform the associated action.
	NoncurrentDays int `json:"noncurrentDays,omitempty" yaml:"noncurrentDays,omitempty"`

	// Class of storage used to store the object. Valid Values: `GLACIER`, `STANDARD_IA`, `ONEZONE_IA`, `INTELLIGENT_TIERING`, `DEEP_ARCHIVE`, `GLACIER_IR`.
	StorageClass string `json:"storageClass,omitempty" yaml:"storageClass,omitempty"`
}
