package types

type Sagemaker_ModelContainerModelDataSourceS3DataSource struct {
	// How the model data is prepared. Allowed values are: `None` and `Gzip`.
	CompressionType string `json:"compressionType,omitempty" yaml:"compressionType,omitempty"`

	// The type of model data to deploy. Allowed values are: `S3Object` and `S3Prefix`.
	S3DataType string `json:"s3DataType,omitempty" yaml:"s3DataType,omitempty"`

	// The S3 path of model data to deploy.
	S3Uri string `json:"s3Uri,omitempty" yaml:"s3Uri,omitempty"`
}
