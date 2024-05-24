package types

type Glue_getDataCatalogEncryptionSettingsDataCatalogEncryptionSetting struct {
	// Encryption-at-rest configuration for the Data Catalog. see Encryption At Rest.
	EncryptionAtRests []Glue_getDataCatalogEncryptionSettingsDataCatalogEncryptionSettingEncryptionAtRest `json:"encryptionAtRests,omitempty" yaml:"encryptionAtRests,omitempty"`

	// When connection password protection is enabled, the Data Catalog uses a customer-provided key to encrypt the password as part of CreateConnection or UpdateConnection and store it in the ENCRYPTED_PASSWORD field in the connection properties. You can enable catalog encryption or only password encryption. see Connection Password Encryption.
	ConnectionPasswordEncryptions []Glue_getDataCatalogEncryptionSettingsDataCatalogEncryptionSettingConnectionPasswordEncryption `json:"connectionPasswordEncryptions,omitempty" yaml:"connectionPasswordEncryptions,omitempty"`
}
