package kms

type KeyRingImportJob struct {
	/*
	   The protection level of the ImportJob. This must match the protectionLevel of the
	   versionTemplate on the CryptoKey you attempt to import into.
	   Possible values are: `SOFTWARE`, `HSM`, `EXTERNAL`.
	*/
	ProtectionLevel string `json:"protectionLevel,omitempty" yaml:"protectionLevel,omitempty"`

	/*
	   It must be unique within a KeyRing and match the regular expression [a-zA-Z0-9_-]{1,63}


	   - - -
	*/
	ImportJobId string `json:"importJobId,omitempty" yaml:"importJobId,omitempty"`

	/*
	   The wrapping method to be used for incoming key material.
	   Possible values are: `RSA_OAEP_3072_SHA1_AES_256`, `RSA_OAEP_4096_SHA1_AES_256`.
	*/
	ImportMethod string `json:"importMethod,omitempty" yaml:"importMethod,omitempty"`

	/*
	   The KeyRing that this import job belongs to.
	   Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
	*/
	KeyRing string `json:"keyRing,omitempty" yaml:"keyRing,omitempty"`
}
