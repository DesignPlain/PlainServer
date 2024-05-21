package types

type Kms_CryptoKeyVersionAttestation struct {
	/*
	   The certificate chains needed to validate the attestation
	   Structure is documented below.
	*/
	CertChains Kms_CryptoKeyVersionAttestationCertChains `json:"certChains,omitempty" yaml:"certChains,omitempty"`

	/*
	   (Output)
	   The attestation data provided by the HSM when the key operation was performed.
	*/
	Content string `json:"content,omitempty" yaml:"content,omitempty"`

	/*
	   ExternalProtectionLevelOptions stores a group of additional fields for configuring a CryptoKeyVersion that are specific to the EXTERNAL protection level and EXTERNAL_VPC protection levels.
	   Structure is documented below.
	*/
	ExternalProtectionLevelOptions Kms_CryptoKeyVersionAttestationExternalProtectionLevelOptions `json:"externalProtectionLevelOptions,omitempty" yaml:"externalProtectionLevelOptions,omitempty"`

	/*
	   (Output)
	   The format of the attestation data.
	*/
	Format string `json:"format,omitempty" yaml:"format,omitempty"`
}
