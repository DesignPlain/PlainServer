package types

type Glue_DataCatalogEncryptionSettingsDataCatalogEncryptionSettings struct {
	// When connection password protection is enabled, the Data Catalog uses a customer-provided key to encrypt the password as part of CreateConnection or UpdateConnection and store it in the ENCRYPTED_PASSWORD field in the connection properties. You can enable catalog encryption or only password encryption. see Connection Password Encryption.
	ConnectionPasswordEncryption Glue_DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryption `json:"connectionPasswordEncryption,omitempty" yaml:"connectionPasswordEncryption,omitempty"`

	// Specifies the encryption-at-rest configuration for the Data Catalog. see Encryption At Rest.
	EncryptionAtRest Glue_DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRest `json:"encryptionAtRest,omitempty" yaml:"encryptionAtRest,omitempty"`
}
