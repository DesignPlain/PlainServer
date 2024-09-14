package s3

import types "libds/aws/types"

type BucketLoggingV2 struct {
	// Name of the bucket.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Account ID of the expected bucket owner.
	ExpectedBucketOwner string `json:"expectedBucketOwner,omitempty" yaml:"expectedBucketOwner,omitempty"`

	// Name of the bucket where you want Amazon S3 to store server access logs.
	TargetBucket string `json:"targetBucket,omitempty" yaml:"targetBucket,omitempty"`

	// Set of configuration blocks with information for granting permissions. See below.
	TargetGrants []types.S3_BucketLoggingV2TargetGrant `json:"targetGrants,omitempty" yaml:"targetGrants,omitempty"`

	// Amazon S3 key format for log objects. See below.
	TargetObjectKeyFormat types.S3_BucketLoggingV2TargetObjectKeyFormat `json:"targetObjectKeyFormat,omitempty" yaml:"targetObjectKeyFormat,omitempty"`

	// Prefix for all log object keys.
	TargetPrefix string `json:"targetPrefix,omitempty" yaml:"targetPrefix,omitempty"`
}
