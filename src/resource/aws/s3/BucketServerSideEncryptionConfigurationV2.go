package s3

import types "DesignSphere_Server/src/resource/aws/types"

type BucketServerSideEncryptionConfigurationV2 struct {
	// ID (name) of the bucket.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Account ID of the expected bucket owner.
	ExpectedBucketOwner string `json:"expectedBucketOwner,omitempty" yaml:"expectedBucketOwner,omitempty"`

	// Set of server-side encryption configuration rules. See below. Currently, only a single rule is supported.
	Rules []types.S3_BucketServerSideEncryptionConfigurationV2Rule `json:"rules,omitempty" yaml:"rules,omitempty"`
}
