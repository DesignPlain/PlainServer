package types

type Kms_CryptoKeyPrimary struct {
	// The resource name for the CryptoKey.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   (Output)
	   The current state of the CryptoKeyVersion.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
