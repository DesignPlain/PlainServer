package types

type Glue_DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest struct {
	// The encryption-at-rest mode for encrypting Data Catalog data. Valid values are `DISABLED` and `SSE-KMS`.
	CatalogEncryptionMode string `json:"catalogEncryptionMode,omitempty" yaml:"catalogEncryptionMode,omitempty"`

	// The ARN of the AWS KMS key to use for encryption at rest.
	SseAwsKmsKeyId string `json:"sseAwsKmsKeyId,omitempty" yaml:"sseAwsKmsKeyId,omitempty"`
}
