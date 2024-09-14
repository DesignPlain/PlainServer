package types

type Kms_getKMSCryptoKeyVersionPublicKey struct {
	// The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
	Algorithm string `json:"algorithm,omitempty" yaml:"algorithm,omitempty"`

	// The public key, encoded in PEM format. For more information, see the RFC 7468 sections for General Considerations and Textual Encoding of Subject Public Key Info.
	Pem string `json:"pem,omitempty" yaml:"pem,omitempty"`
}
