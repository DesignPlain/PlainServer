package s3

import types "DesignSphere_Server/src/resource/aws/types"

type BucketLifecycleConfigurationV2 struct {
	// Name of the source S3 bucket you want Amazon S3 to monitor.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner string `json:"expectedBucketOwner,omitempty" yaml:"expectedBucketOwner,omitempty"`

	// List of configuration blocks describing the rules managing the replication. See below.
	Rules []types.S3_BucketLifecycleConfigurationV2Rule `json:"rules,omitempty" yaml:"rules,omitempty"`
}
