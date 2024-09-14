package types

type Glue_DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest struct {
	// The encryption-at-rest mode for encrypting Data Catalog data. Valid values: `DISABLED`, `SSE-KMS`, `SSE-KMS-WITH-SERVICE-ROLE`.
	CatalogEncryptionMode string `json:"catalogEncryptionMode,omitempty" yaml:"catalogEncryptionMode,omitempty"`

	// The ARN of the AWS IAM role used for accessing encrypted Data Catalog data.
	CatalogEncryptionServiceRole string `json:"catalogEncryptionServiceRole,omitempty" yaml:"catalogEncryptionServiceRole,omitempty"`

	// The ARN of the AWS KMS key to use for encryption at rest.
	SseAwsKmsKeyId string `json:"sseAwsKmsKeyId,omitempty" yaml:"sseAwsKmsKeyId,omitempty"`
}
