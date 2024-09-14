package types

type Imagebuilder_InfrastructureConfigurationLoggingS3Logs struct {
	/*
	   Name of the S3 Bucket.

	   The following arguments are optional:
	*/
	S3BucketName string `json:"s3BucketName,omitempty" yaml:"s3BucketName,omitempty"`

	// Prefix to use for S3 logs. Defaults to `/`.
	S3KeyPrefix string `json:"s3KeyPrefix,omitempty" yaml:"s3KeyPrefix,omitempty"`
}
