package types

type Sagemaker_PipelinePipelineDefinitionS3Location struct {
	// Name of the S3 bucket.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// The object key (or key name) uniquely identifies the object in an S3 bucket.
	ObjectKey string `json:"objectKey,omitempty" yaml:"objectKey,omitempty"`

	// Version Id of the pipeline definition file. If not specified, Amazon SageMaker will retrieve the latest version.
	VersionId string `json:"versionId,omitempty" yaml:"versionId,omitempty"`
}
