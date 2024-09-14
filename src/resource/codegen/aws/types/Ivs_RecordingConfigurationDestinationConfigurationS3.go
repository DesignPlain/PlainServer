package types

type Ivs_RecordingConfigurationDestinationConfigurationS3 struct {
	/*
	   S3 bucket name where recorded videos will be stored.

	   The following arguments are optional:
	*/
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`
}
