package types

type Paymentcryptography_KeyKeyAttributes struct {
	// Key algorithm to be use during creation of an AWS Payment Cryptography key.
	KeyAlgorithm string `json:"keyAlgorithm,omitempty" yaml:"keyAlgorithm,omitempty"`

	// Type of AWS Payment Cryptography key to create.
	KeyClass string `json:"keyClass,omitempty" yaml:"keyClass,omitempty"`

	// List of cryptographic operations that you can perform using the key.
	KeyModesOfUse Paymentcryptography_KeyKeyAttributesKeyModesOfUse `json:"keyModesOfUse,omitempty" yaml:"keyModesOfUse,omitempty"`

	// Cryptographic usage of an AWS Payment Cryptography key as defined in section A.5.2 of the TR-31 spec.
	KeyUsage string `json:"keyUsage,omitempty" yaml:"keyUsage,omitempty"`
}
