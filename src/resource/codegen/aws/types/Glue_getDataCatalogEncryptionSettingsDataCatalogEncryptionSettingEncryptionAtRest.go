package types

type Glue_getDataCatalogEncryptionSettingsDataCatalogEncryptionSettingEncryptionAtRest struct {
	// The encryption-at-rest mode for encrypting Data Catalog data.
	CatalogEncryptionMode string `json:"catalogEncryptionMode,omitempty" yaml:"catalogEncryptionMode,omitempty"`

	// The ARN of the AWS IAM role used for accessing encrypted Data Catalog data.
	CatalogEncryptionServiceRole string `json:"catalogEncryptionServiceRole,omitempty" yaml:"catalogEncryptionServiceRole,omitempty"`

	// ARN of the AWS KMS key to use for encryption at rest.
	SseAwsKmsKeyId string `json:"sseAwsKmsKeyId,omitempty" yaml:"sseAwsKmsKeyId,omitempty"`
}
