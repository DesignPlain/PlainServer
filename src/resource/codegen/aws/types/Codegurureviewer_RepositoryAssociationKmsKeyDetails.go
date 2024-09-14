package types

type Codegurureviewer_RepositoryAssociationKmsKeyDetails struct {
	// The encryption option for a repository association. It is either owned by AWS Key Management Service (KMS) (`AWS_OWNED_CMK`) or customer managed (`CUSTOMER_MANAGED_CMK`).
	EncryptionOption string `json:"encryptionOption,omitempty" yaml:"encryptionOption,omitempty"`

	// The ID of the AWS KMS key that is associated with a repository association.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`
}
