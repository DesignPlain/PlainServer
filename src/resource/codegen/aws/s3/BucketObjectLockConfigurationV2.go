package s3

import types "libds/aws/types"

type BucketObjectLockConfigurationV2 struct {
	// Name of the bucket.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Account ID of the expected bucket owner.
	ExpectedBucketOwner string `json:"expectedBucketOwner,omitempty" yaml:"expectedBucketOwner,omitempty"`

	// Indicates whether this bucket has an Object Lock configuration enabled. Defaults to `Enabled`. Valid values: `Enabled`.
	ObjectLockEnabled string `json:"objectLockEnabled,omitempty" yaml:"objectLockEnabled,omitempty"`

	// Configuration block for specifying the Object Lock rule for the specified object. See below.
	Rule types.S3_BucketObjectLockConfigurationV2Rule `json:"rule,omitempty" yaml:"rule,omitempty"`

	/*
	   This argument is deprecated and no longer needed to enable Object Lock.
	   To enable Object Lock for an existing bucket, you must first enable versioning on the bucket and then enable Object Lock. For more details on versioning, see the `aws.s3.BucketVersioningV2` resource.
	*/
	Token string `json:"token,omitempty" yaml:"token,omitempty"`
}
