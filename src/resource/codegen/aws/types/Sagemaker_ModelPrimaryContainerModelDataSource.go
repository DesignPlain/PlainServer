package types

type Sagemaker_ModelPrimaryContainerModelDataSource struct {
	// The S3 location of model data to deploy.
	S3DataSources []Sagemaker_ModelPrimaryContainerModelDataSourceS3DataSource `json:"s3DataSources,omitempty" yaml:"s3DataSources,omitempty"`
}
