package types

type Verifiedaccess_InstanceLoggingConfigurationAccessLogsS3 struct {
	// The name of S3 bucket.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// The ID of the AWS account that owns the Amazon S3 bucket.
	BucketOwner string `json:"bucketOwner,omitempty" yaml:"bucketOwner,omitempty"`

	// Indicates whether logging is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The bucket prefix.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
}
