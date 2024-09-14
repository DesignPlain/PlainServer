package types

type Certificateauthority_CertificateConfigX509ConfigKeyUsageBaseKeyUsage struct {
	// The key may be used for cryptographic commitments. Note that this may also be referred to as "non-repudiation".
	ContentCommitment bool `json:"contentCommitment,omitempty" yaml:"contentCommitment,omitempty"`

	// The key may be used to decipher only.
	DecipherOnly bool `json:"decipherOnly,omitempty" yaml:"decipherOnly,omitempty"`

	// The key may be used to encipher only.
	EncipherOnly bool `json:"encipherOnly,omitempty" yaml:"encipherOnly,omitempty"`

	// The key may be used in a key agreement protocol.
	KeyAgreement bool `json:"keyAgreement,omitempty" yaml:"keyAgreement,omitempty"`

	// The key may be used to encipher other keys.
	KeyEncipherment bool `json:"keyEncipherment,omitempty" yaml:"keyEncipherment,omitempty"`

	// The key may be used to sign certificates.
	CertSign bool `json:"certSign,omitempty" yaml:"certSign,omitempty"`

	// The key may be used sign certificate revocation lists.
	CrlSign bool `json:"crlSign,omitempty" yaml:"crlSign,omitempty"`

	// The key may be used to encipher data.
	DataEncipherment bool `json:"dataEncipherment,omitempty" yaml:"dataEncipherment,omitempty"`

	// The key may be used for digital signatures.
	DigitalSignature bool `json:"digitalSignature,omitempty" yaml:"digitalSignature,omitempty"`
}
