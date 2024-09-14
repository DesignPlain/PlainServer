package types

type Opensearchingest_PipelineEncryptionAtRestOptions struct {
	// The ARN of the KMS key used to encrypt data-at-rest in OpenSearch Ingestion. By default, data is encrypted using an AWS owned key.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`
}
