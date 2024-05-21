package types

type Kms_KeyRingImportJobPublicKey struct {
	/*
	   (Output)
	   The public key, encoded in PEM format. For more information, see the RFC 7468 sections
	   for General Considerations and Textual Encoding of Subject Public Key Info.
	*/
	Pem string `json:"pem,omitempty" yaml:"pem,omitempty"`
}
