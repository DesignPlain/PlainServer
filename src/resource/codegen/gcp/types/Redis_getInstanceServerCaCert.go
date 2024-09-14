package types

type Redis_getInstanceServerCaCert struct {
	// The time when the certificate was created.
	CreateTime string `json:"createTime,omitempty" yaml:"createTime,omitempty"`

	// The time when the certificate expires.
	ExpireTime string `json:"expireTime,omitempty" yaml:"expireTime,omitempty"`

	// Serial number, as extracted from the certificate.
	SerialNumber string `json:"serialNumber,omitempty" yaml:"serialNumber,omitempty"`

	// Sha1 Fingerprint of the certificate.
	Sha1Fingerprint string `json:"sha1Fingerprint,omitempty" yaml:"sha1Fingerprint,omitempty"`

	// The certificate data in PEM format.
	Cert string `json:"cert,omitempty" yaml:"cert,omitempty"`
}
