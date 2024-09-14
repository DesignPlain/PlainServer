package types

type Certificateauthority_CaPoolIssuancePolicyAllowedKeyTypeRsa struct {
	/*
	   The maximum allowed RSA modulus size, in bits. If this is not set, or if set to zero, the
	   service will not enforce an explicit upper bound on RSA modulus sizes.
	*/
	MaxModulusSize string `json:"maxModulusSize,omitempty" yaml:"maxModulusSize,omitempty"`

	/*
	   The minimum allowed RSA modulus size, in bits. If this is not set, or if set to zero, the
	   service-level min RSA modulus size will continue to apply.
	*/
	MinModulusSize string `json:"minModulusSize,omitempty" yaml:"minModulusSize,omitempty"`
}
