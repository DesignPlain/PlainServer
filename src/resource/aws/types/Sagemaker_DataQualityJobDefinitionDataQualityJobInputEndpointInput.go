package types

type Sagemaker_DataQualityJobDefinitionDataQualityJobInputEndpointInput struct {
	// Whether the `Pipe` or `File` is used as the input mode for transferring data for the monitoring job. `Pipe` mode is recommended for large datasets. `File` mode is useful for small files that fit in memory. Defaults to `File`.  Valid values are `Pipe` or `File`
	S3InputMode string `json:"s3InputMode,omitempty" yaml:"s3InputMode,omitempty"`

	// An endpoint in customer's account which has `data_capture_config` enabled.
	EndpointName string `json:"endpointName,omitempty" yaml:"endpointName,omitempty"`

	// Path to the filesystem where the endpoint data is available to the container. Defaults to `/opt/ml/processing/input`.
	LocalPath string `json:"localPath,omitempty" yaml:"localPath,omitempty"`

	// Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defaults to `FullyReplicated`. Valid values are `FullyReplicated` or `ShardedByS3Key`
	S3DataDistributionType string `json:"s3DataDistributionType,omitempty" yaml:"s3DataDistributionType,omitempty"`
}
