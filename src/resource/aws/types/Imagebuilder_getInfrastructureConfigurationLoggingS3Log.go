package types

type Imagebuilder_getInfrastructureConfigurationLoggingS3Log struct {
	// Name of the S3 Bucket for logging.
	S3BucketName string `json:"s3BucketName,omitempty" yaml:"s3BucketName,omitempty"`

	// Key prefix for S3 Bucket logging.
	S3KeyPrefix string `json:"s3KeyPrefix,omitempty" yaml:"s3KeyPrefix,omitempty"`
}
