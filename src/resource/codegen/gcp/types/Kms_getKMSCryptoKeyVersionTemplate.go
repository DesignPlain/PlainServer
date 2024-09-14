package types

type Kms_getKMSCryptoKeyVersionTemplate struct {
	/*
	   The algorithm to use when creating a version based on this template.
	   See the [algorithm reference](https://cloud.google.com/kms/docs/reference/rest/v1/CryptoKeyVersionAlgorithm) for possible inputs.
	*/
	Algorithm string `json:"algorithm,omitempty" yaml:"algorithm,omitempty"`

	// The protection level to use when creating a version based on this template. Possible values include "SOFTWARE", "HSM", "EXTERNAL", "EXTERNAL_VPC". Defaults to "SOFTWARE".
	ProtectionLevel string `json:"protectionLevel,omitempty" yaml:"protectionLevel,omitempty"`
}
