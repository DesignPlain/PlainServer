package types

type Sagemaker_FlowDefinitionOutputConfig struct {
	// The Amazon Key Management Service (KMS) key ARN for server-side encryption.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// The Amazon S3 path where the object containing human output will be made available.
	S3OutputPath string `json:"s3OutputPath,omitempty" yaml:"s3OutputPath,omitempty"`
}
