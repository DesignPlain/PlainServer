package types

type Binaryauthorization_AttestorAttestationAuthorityNotePublicKeyPkixPublicKey struct {
	/*
	   A PEM-encoded public key, as described in
	   `https://tools.ietf.org/html/rfc7468#section-13`
	*/
	PublicKeyPem string `json:"publicKeyPem,omitempty" yaml:"publicKeyPem,omitempty"`

	/*
	   The signature algorithm used to verify a message against
	   a signature using this key. These signature algorithm must
	   match the structure and any object identifiers encoded in
	   publicKeyPem (i.e. this algorithm must match that of the
	   public key).

	   - - -
	*/
	SignatureAlgorithm string `json:"signatureAlgorithm,omitempty" yaml:"signatureAlgorithm,omitempty"`
}
