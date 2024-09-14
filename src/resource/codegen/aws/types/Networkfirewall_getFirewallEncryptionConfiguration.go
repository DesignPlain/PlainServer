package types

type Networkfirewall_getFirewallEncryptionConfiguration struct {
	// The ID of the AWS Key Management Service (AWS KMS) customer managed key.
	KeyId string `json:"keyId,omitempty" yaml:"keyId,omitempty"`

	// The type of the AWS Key Management Service (AWS KMS) key use by the firewall.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
