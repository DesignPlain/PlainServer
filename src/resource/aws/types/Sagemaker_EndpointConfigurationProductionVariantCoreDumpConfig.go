package types

type Sagemaker_EndpointConfigurationProductionVariantCoreDumpConfig struct {
	// The Amazon S3 bucket to send the core dump to.
	DestinationS3Uri string `json:"destinationS3Uri,omitempty" yaml:"destinationS3Uri,omitempty"`

	// The Amazon Web Services Key Management Service (Amazon Web Services KMS) key that SageMaker uses to encrypt the core dump data at rest using Amazon S3 server-side encryption.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`
}
