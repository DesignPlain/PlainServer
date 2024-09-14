package types

type Pipes_PipeLogConfigurationS3LogDestination struct {
	// Amazon Web Services account that owns the Amazon S3 bucket to which EventBridge delivers the log records for the pipe.
	BucketOwner string `json:"bucketOwner,omitempty" yaml:"bucketOwner,omitempty"`

	// EventBridge format for the log records. Valid values `json`, `plain` and `w3c`.
	OutputFormat string `json:"outputFormat,omitempty" yaml:"outputFormat,omitempty"`

	// Prefix text with which to begin Amazon S3 log object names.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// Name of the Amazon S3 bucket to which EventBridge delivers the log records for the pipe.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`
}
