package types

type Kms_getKMSCryptoKeyPrimary struct {
	/*
	   The CryptoKey's name.
	   A CryptoKey’s name belonging to the specified Google Cloud Platform KeyRing and match the regular expression `[a-zA-Z0-9_-]{1,63}`
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The current state of the CryptoKeyVersion.
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
