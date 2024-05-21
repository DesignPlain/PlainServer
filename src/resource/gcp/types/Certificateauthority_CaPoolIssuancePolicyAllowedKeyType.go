package types



type Certificateauthority_CaPoolIssuancePolicyAllowedKeyType struct {
	/*
	   Represents an allowed Elliptic Curve key type.
	   Structure is documented below.
	*/
	EllipticCurve Certificateauthority_CaPoolIssuancePolicyAllowedKeyTypeEllipticCurve `json:"ellipticCurve,omitempty" yaml:"ellipticCurve,omitempty"`

	/*
	   Describes an RSA key that may be used in a Certificate issued from a CaPool.
	   Structure is documented below.
	*/
	Rsa Certificateauthority_CaPoolIssuancePolicyAllowedKeyTypeRsa `json:"rsa,omitempty" yaml:"rsa,omitempty"`
}
