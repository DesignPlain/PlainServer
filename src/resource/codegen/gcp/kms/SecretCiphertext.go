package kms

type SecretCiphertext struct {
	/*
	   The additional authenticated data used for integrity checks during encryption and decryption.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	AdditionalAuthenticatedData string `json:"additionalAuthenticatedData,omitempty" yaml:"additionalAuthenticatedData,omitempty"`

	/*
	   The full name of the CryptoKey that will be used to encrypt the provided plaintext.
	   Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}'`


	   - - -
	*/
	CryptoKey string `json:"cryptoKey,omitempty" yaml:"cryptoKey,omitempty"`

	/*
	   The plaintext to be encrypted.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	Plaintext string `json:"plaintext,omitempty" yaml:"plaintext,omitempty"`
}
