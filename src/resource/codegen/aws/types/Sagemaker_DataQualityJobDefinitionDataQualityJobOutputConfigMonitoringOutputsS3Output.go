package types

type Sagemaker_DataQualityJobDefinitionDataQualityJobOutputConfigMonitoringOutputsS3Output struct {
	// The local path to the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job. LocalPath is an absolute path for the output data. Defaults to `/opt/ml/processing/output`.
	LocalPath string `json:"localPath,omitempty" yaml:"localPath,omitempty"`

	// Whether to upload the results of the monitoring job continuously or after the job completes. Valid values are `Continuous` or `EndOfJob`
	S3UploadMode string `json:"s3UploadMode,omitempty" yaml:"s3UploadMode,omitempty"`

	// A URI that identifies the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job.
	S3Uri string `json:"s3Uri,omitempty" yaml:"s3Uri,omitempty"`
}
