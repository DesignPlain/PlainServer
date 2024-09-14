package types

type Keyspaces_TableEncryptionSpecification struct {
	// The Amazon Resource Name (ARN) of the customer managed KMS key.
	KmsKeyIdentifier string `json:"kmsKeyIdentifier,omitempty" yaml:"kmsKeyIdentifier,omitempty"`

	// The encryption option specified for the table. Valid values: `AWS_OWNED_KMS_KEY`, `CUSTOMER_MANAGED_KMS_KEY`. The default value is `AWS_OWNED_KMS_KEY`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
