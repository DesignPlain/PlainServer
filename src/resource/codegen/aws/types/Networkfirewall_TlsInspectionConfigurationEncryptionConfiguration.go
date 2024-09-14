package types

type Networkfirewall_TlsInspectionConfigurationEncryptionConfiguration struct {
	// ARN of the Amazon Web Services Key Management Service (KMS) customer managed key.
	KeyId string `json:"keyId,omitempty" yaml:"keyId,omitempty"`

	// Type of KMS key to use for encryption of your Network Firewall resources. Valid values: `AWS_OWNED_KMS_KEY`, `CUSTOMER_KMS`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
