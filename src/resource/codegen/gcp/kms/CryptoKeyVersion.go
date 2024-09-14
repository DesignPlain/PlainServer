package kms

type CryptoKeyVersion struct {
	/*
	   The name of the cryptoKey associated with the CryptoKeyVersions.
	   Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}'`


	   - - -
	*/
	CryptoKey string `json:"cryptoKey,omitempty" yaml:"cryptoKey,omitempty"`

	/*
	   The current state of the CryptoKeyVersion.
	   Possible values are: `PENDING_GENERATION`, `ENABLED`, `DISABLED`, `DESTROYED`, `DESTROY_SCHEDULED`, `PENDING_IMPORT`, `IMPORT_FAILED`.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
