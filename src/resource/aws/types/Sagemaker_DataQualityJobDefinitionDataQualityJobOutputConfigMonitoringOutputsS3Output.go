package types

type Sagemaker_DataQualityJobDefinitionDataQualityJobOutputConfigMonitoringOutputsS3Output struct {
	// Path to the filesystem where the batch transform data is available to the container. Defaults to `/opt/ml/processing/input`.
	LocalPath string `json:"localPath,omitempty" yaml:"localPath,omitempty"`

	// Whether to upload the results of the monitoring job continuously or after the job completes. Valid values are `Continuous` or `EndOfJob`
	S3UploadMode string `json:"s3UploadMode,omitempty" yaml:"s3UploadMode,omitempty"`

	// The Amazon S3 URI for the constraints resource.
	S3Uri string `json:"s3Uri,omitempty" yaml:"s3Uri,omitempty"`
}
