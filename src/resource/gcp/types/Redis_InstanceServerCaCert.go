package types

type Redis_InstanceServerCaCert struct {
	/*
	   (Output)
	   The certificate data in PEM format.
	*/
	Cert string `json:"cert,omitempty" yaml:"cert,omitempty"`

	/*
	   (Output)
	   Output only. The time when the policy was created.
	   A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	   resolution and up to nine fractional digits.
	*/
	CreateTime string `json:"createTime,omitempty" yaml:"createTime,omitempty"`

	/*
	   (Output)
	   The time when the certificate expires.
	*/
	ExpireTime string `json:"expireTime,omitempty" yaml:"expireTime,omitempty"`

	/*
	   (Output)
	   Serial number, as extracted from the certificate.
	*/
	SerialNumber string `json:"serialNumber,omitempty" yaml:"serialNumber,omitempty"`

	/*
	   (Output)
	   Sha1 Fingerprint of the certificate.
	*/
	Sha1Fingerprint string `json:"sha1Fingerprint,omitempty" yaml:"sha1Fingerprint,omitempty"`
}
