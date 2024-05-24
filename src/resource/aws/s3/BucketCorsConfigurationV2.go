package s3

import types "DesignSphere_Server/src/resource/aws/types"

type BucketCorsConfigurationV2 struct {
	// Account ID of the expected bucket owner.
	ExpectedBucketOwner string `json:"expectedBucketOwner,omitempty" yaml:"expectedBucketOwner,omitempty"`

	// Name of the bucket.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Set of origins and methods (cross-origin access that you want to allow). See below. You can configure up to 100 rules.
	CorsRules []types.S3_BucketCorsConfigurationV2CorsRule `json:"corsRules,omitempty" yaml:"corsRules,omitempty"`
}
