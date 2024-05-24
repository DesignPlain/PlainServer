package s3

import types "DesignSphere_Server/src/resource/aws/types"

type BucketMetric struct {
	// Name of the bucket to put metric configuration.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
	Filter types.S3_BucketMetricFilter `json:"filter,omitempty" yaml:"filter,omitempty"`

	// Unique identifier of the metrics configuration for the bucket. Must be less than or equal to 64 characters in length.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
