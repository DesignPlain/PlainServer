package types

type Paymentcryptography_KeyKeyAttributesKeyModesOfUse struct {
	// Whether an AWS Payment Cryptography key can be used to encrypt data.
	Encrypt bool `json:"encrypt,omitempty" yaml:"encrypt,omitempty"`

	// Whether an AWS Payment Cryptography key can be used for signing.
	Sign bool `json:"sign,omitempty" yaml:"sign,omitempty"`

	// Whether an AWS Payment Cryptography key can be used to verify signatures.
	Verify bool `json:"verify,omitempty" yaml:"verify,omitempty"`

	// Whether an AWS Payment Cryptography key can be used to wrap other keys.
	Wrap bool `json:"wrap,omitempty" yaml:"wrap,omitempty"`

	// Whether an AWS Payment Cryptography key can be used to decrypt data.
	Decrypt bool `json:"decrypt,omitempty" yaml:"decrypt,omitempty"`

	// Whether an AWS Payment Cryptography key can be used to derive new keys.
	DeriveKey bool `json:"deriveKey,omitempty" yaml:"deriveKey,omitempty"`

	// Whether an AWS Payment Cryptography key can be used to generate and verify other card and PIN verification keys.
	Generate bool `json:"generate,omitempty" yaml:"generate,omitempty"`

	// Whether an AWS Payment Cryptography key has no special restrictions other than the restrictions implied by KeyUsage.
	NoRestrictions bool `json:"noRestrictions,omitempty" yaml:"noRestrictions,omitempty"`

	// Whether an AWS Payment Cryptography key can be used to unwrap other keys.
	Unwrap bool `json:"unwrap,omitempty" yaml:"unwrap,omitempty"`
}
