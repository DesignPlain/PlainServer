package types

type Kms_KeyRingImportJobAttestation struct {
	/*
	   (Output)
	   The attestation data provided by the HSM when the key operation was performed.
	   A base64-encoded string.
	*/
	Content string `json:"content,omitempty" yaml:"content,omitempty"`

	/*
	   (Output)
	   The format of the attestation data.
	*/
	Format string `json:"format,omitempty" yaml:"format,omitempty"`
}
