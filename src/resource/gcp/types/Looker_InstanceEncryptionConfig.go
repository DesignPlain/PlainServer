package types

type Looker_InstanceEncryptionConfig struct {
	// Name of the customer managed encryption key (CMEK) in KMS.
	KmsKeyName string `json:"kmsKeyName,omitempty" yaml:"kmsKeyName,omitempty"`

	/*
	   (Output)
	   Full name and version of the CMEK key currently in use to encrypt Looker data.
	*/
	KmsKeyNameVersion string `json:"kmsKeyNameVersion,omitempty" yaml:"kmsKeyNameVersion,omitempty"`

	/*
	   (Output)
	   Status of the customer managed encryption key (CMEK) in KMS.
	*/
	KmsKeyState string `json:"kmsKeyState,omitempty" yaml:"kmsKeyState,omitempty"`
}
