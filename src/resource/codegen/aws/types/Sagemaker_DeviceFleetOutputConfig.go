package types

type Sagemaker_DeviceFleetOutputConfig struct {
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume after compilation job. If you don't provide a KMS key ID, Amazon SageMaker uses the default KMS key for Amazon S3 for your role's account.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// The Amazon Simple Storage (S3) bucker URI.
	S3OutputLocation string `json:"s3OutputLocation,omitempty" yaml:"s3OutputLocation,omitempty"`
}
