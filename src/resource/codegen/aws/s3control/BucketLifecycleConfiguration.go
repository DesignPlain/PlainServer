package s3control

import types "libds/aws/types"

type BucketLifecycleConfiguration struct {
	// Amazon Resource Name (ARN) of the bucket.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Configuration block(s) containing lifecycle rules for the bucket.
	Rules []types.S3control_BucketLifecycleConfigurationRule `json:"rules,omitempty" yaml:"rules,omitempty"`
}
