package types

type Securitylake_DataLakeConfigurationEncryptionConfiguration struct {
	// The id of KMS encryption key used by Amazon Security Lake to encrypt the Security Lake object.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`
}
