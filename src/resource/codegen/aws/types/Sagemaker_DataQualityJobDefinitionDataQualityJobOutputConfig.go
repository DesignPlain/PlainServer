package types

type Sagemaker_DataQualityJobDefinitionDataQualityJobOutputConfig struct {
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Monitoring outputs for monitoring jobs. This is where the output of the periodic monitoring jobs is uploaded. Fields are documented below.
	MonitoringOutputs Sagemaker_DataQualityJobDefinitionDataQualityJobOutputConfigMonitoringOutputs `json:"monitoringOutputs,omitempty" yaml:"monitoringOutputs,omitempty"`
}
