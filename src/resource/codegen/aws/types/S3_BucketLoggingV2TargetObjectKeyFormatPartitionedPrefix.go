package types

type S3_BucketLoggingV2TargetObjectKeyFormatPartitionedPrefix struct {
	// Specifies the partition date source for the partitioned prefix. Valid values: `EventTime`, `DeliveryTime`.
	PartitionDateSource string `json:"partitionDateSource,omitempty" yaml:"partitionDateSource,omitempty"`
}
