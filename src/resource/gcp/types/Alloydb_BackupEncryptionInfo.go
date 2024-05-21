package types

type Alloydb_BackupEncryptionInfo struct {
	/*
	   (Output)
	   Output only. Type of encryption.
	*/
	EncryptionType string `json:"encryptionType,omitempty" yaml:"encryptionType,omitempty"`

	/*
	   (Output)
	   Output only. Cloud KMS key versions that are being used to protect the database or the backup.
	*/
	KmsKeyVersions []string `json:"kmsKeyVersions,omitempty" yaml:"kmsKeyVersions,omitempty"`
}
