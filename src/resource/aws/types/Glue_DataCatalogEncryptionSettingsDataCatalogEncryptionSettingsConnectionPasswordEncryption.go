package types

type Glue_DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption struct {
	// A KMS key ARN that is used to encrypt the connection password. If connection password protection is enabled, the caller of CreateConnection and UpdateConnection needs at least `kms:Encrypt` permission on the specified AWS KMS key, to encrypt passwords before storing them in the Data Catalog.
	AwsKmsKeyId string `json:"awsKmsKeyId,omitempty" yaml:"awsKmsKeyId,omitempty"`

	// When set to `true`, passwords remain encrypted in the responses of GetConnection and GetConnections. This encryption takes effect independently of the catalog encryption.
	ReturnConnectionPasswordEncrypted bool `json:"returnConnectionPasswordEncrypted,omitempty" yaml:"returnConnectionPasswordEncrypted,omitempty"`
}
