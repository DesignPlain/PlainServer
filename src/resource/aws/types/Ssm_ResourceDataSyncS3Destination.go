package types

type Ssm_ResourceDataSyncS3Destination struct {
	// ARN of an encryption key for a destination in Amazon S3.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	// Prefix for the bucket.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// Region with the bucket targeted by the Resource Data Sync.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// A supported sync format. Only JsonSerDe is currently supported. Defaults to JsonSerDe.
	SyncFormat string `json:"syncFormat,omitempty" yaml:"syncFormat,omitempty"`

	// Name of S3 bucket where the aggregated data is stored.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`
}
