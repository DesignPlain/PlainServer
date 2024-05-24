package types

type S3_BucketLoggingV2TargetObjectKeyFormat struct {
	// Partitioned S3 key for log objects. See below.
	PartitionedPrefix S3_BucketLoggingV2TargetObjectKeyFormatPartitionedPrefix `json:"partitionedPrefix,omitempty" yaml:"partitionedPrefix,omitempty"`

	// Use the simple format for S3 keys for log objects. To use, set `simple_prefix {}`.
	SimplePrefix S3_BucketLoggingV2TargetObjectKeyFormatSimplePrefix `json:"simplePrefix,omitempty" yaml:"simplePrefix,omitempty"`
}
