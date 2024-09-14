package types

type Sagemaker_DataQualityJobDefinitionDataQualityJobInputBatchTransformInput struct {
	// Whether the `Pipe` or `File` is used as the input mode for transferring data for the monitoring job. `Pipe` mode is recommended for large datasets. `File` mode is useful for small files that fit in memory. Defaults to `File`.  Valid values are `Pipe` or `File`
	S3InputMode string `json:"s3InputMode,omitempty" yaml:"s3InputMode,omitempty"`

	// The Amazon S3 location being used to capture the data.
	DataCapturedDestinationS3Uri string `json:"dataCapturedDestinationS3Uri,omitempty" yaml:"dataCapturedDestinationS3Uri,omitempty"`

	// The dataset format for your batch transform job. Fields are documented below.
	DatasetFormat Sagemaker_DataQualityJobDefinitionDataQualityJobInputBatchTransformInputDatasetFormat `json:"datasetFormat,omitempty" yaml:"datasetFormat,omitempty"`

	// Path to the filesystem where the batch transform data is available to the container. Defaults to `/opt/ml/processing/input`.
	LocalPath string `json:"localPath,omitempty" yaml:"localPath,omitempty"`

	// Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defaults to `FullyReplicated`. Valid values are `FullyReplicated` or `ShardedByS3Key`
	S3DataDistributionType string `json:"s3DataDistributionType,omitempty" yaml:"s3DataDistributionType,omitempty"`
}
