package types

type Kms_CryptoKeyVersionAttestationCertChains struct {
	// Cavium certificate chain corresponding to the attestation.
	CaviumCerts string `json:"caviumCerts,omitempty" yaml:"caviumCerts,omitempty"`

	// Google card certificate chain corresponding to the attestation.
	GoogleCardCerts string `json:"googleCardCerts,omitempty" yaml:"googleCardCerts,omitempty"`

	// Google partition certificate chain corresponding to the attestation.
	GooglePartitionCerts string `json:"googlePartitionCerts,omitempty" yaml:"googlePartitionCerts,omitempty"`
}
