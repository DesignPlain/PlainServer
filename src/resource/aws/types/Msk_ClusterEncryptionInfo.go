package types

type Msk_ClusterEncryptionInfo struct {
	// You may specify a KMS key short ID or ARN (it will always output an ARN) to use for encrypting your data at rest.  If no key is specified, an AWS managed KMS ('aws/msk' managed service) key will be used for encrypting the data at rest.
	EncryptionAtRestKmsKeyArn string `json:"encryptionAtRestKmsKeyArn,omitempty" yaml:"encryptionAtRestKmsKeyArn,omitempty"`

	// Configuration block to specify encryption in transit. See below.
	EncryptionInTransit Msk_ClusterEncryptionInfoEncryptionInTransit `json:"encryptionInTransit,omitempty" yaml:"encryptionInTransit,omitempty"`
}
