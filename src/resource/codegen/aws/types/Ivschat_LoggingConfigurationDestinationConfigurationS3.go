package types

type Ivschat_LoggingConfigurationDestinationConfigurationS3 struct {
	/*
	   Name of the Amazon S3 bucket where chat activity will be logged.

	   The following arguments are optional:
	*/
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`
}
