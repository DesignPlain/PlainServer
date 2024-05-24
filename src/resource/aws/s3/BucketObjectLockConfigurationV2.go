package s3

import types "DesignSphere_Server/src/resource/aws/types"

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
	   Token to allow Object Lock to be enabled for an existing bucket. You must contact AWS support for the bucket's "Object Lock token".
	   The token is generated in the back-end when [versioning](https://docs.aws.amazon.com/AmazonS3/latest/userguide/manage-versioning-examples.html) is enabled on a bucket. For more details on versioning, see the `aws.s3.BucketVersioningV2` resource.
	*/
	Token string `json:"token,omitempty" yaml:"token,omitempty"`
}
