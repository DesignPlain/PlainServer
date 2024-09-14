package types

type Gkebackup_BackupPlanBackupConfigEncryptionKey struct {
	// Google Cloud KMS encryption key. Format: projects/-/locations/-/keyRings/-/cryptoKeys/-
	GcpKmsEncryptionKey string `json:"gcpKmsEncryptionKey,omitempty" yaml:"gcpKmsEncryptionKey,omitempty"`
}
