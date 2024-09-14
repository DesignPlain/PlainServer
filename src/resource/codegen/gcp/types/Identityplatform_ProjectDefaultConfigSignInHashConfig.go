package types

type Identityplatform_ProjectDefaultConfigSignInHashConfig struct {
	/*
	   (Output)
	   Memory cost for hash calculation. Used by scrypt and other similar password derivation algorithms. See https://tools.ietf.org/html/rfc7914 for explanation of field.
	*/
	MemoryCost int `json:"memoryCost,omitempty" yaml:"memoryCost,omitempty"`

	/*
	   (Output)
	   How many rounds for hash calculation. Used by scrypt and other similar password derivation algorithms.
	*/
	Rounds int `json:"rounds,omitempty" yaml:"rounds,omitempty"`

	/*
	   (Output)
	   Non-printable character to be inserted between the salt and plain text password in base64.
	*/
	SaltSeparator string `json:"saltSeparator,omitempty" yaml:"saltSeparator,omitempty"`

	/*
	   (Output)
	   Signer key in base64.
	*/
	SignerKey string `json:"signerKey,omitempty" yaml:"signerKey,omitempty"`

	/*
	   (Output)
	   Different password hash algorithms used in Identity Toolkit.
	*/
	Algorithm string `json:"algorithm,omitempty" yaml:"algorithm,omitempty"`
}
