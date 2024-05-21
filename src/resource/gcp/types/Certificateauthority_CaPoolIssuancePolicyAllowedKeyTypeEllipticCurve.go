package types

type Certificateauthority_CaPoolIssuancePolicyAllowedKeyTypeEllipticCurve struct {
	/*
	   The algorithm used.
	   Possible values are: `ECDSA_P256`, `ECDSA_P384`, `EDDSA_25519`.
	*/
	SignatureAlgorithm string `json:"signatureAlgorithm,omitempty" yaml:"signatureAlgorithm,omitempty"`
}
