package types

type Glue_getDataCatalogEncryptionSettingsDataCatalogEncryptionSettingConnectionPasswordEncryption struct {
	// KMS key ARN that is used to encrypt the connection password.
	AwsKmsKeyId string `json:"awsKmsKeyId,omitempty" yaml:"awsKmsKeyId,omitempty"`

	// When set to `true`, passwords remain encrypted in the responses of GetConnection and GetConnections. This encryption takes effect independently of the catalog encryption.
	ReturnConnectionPasswordEncrypted bool `json:"returnConnectionPasswordEncrypted,omitempty" yaml:"returnConnectionPasswordEncrypted,omitempty"`
}
