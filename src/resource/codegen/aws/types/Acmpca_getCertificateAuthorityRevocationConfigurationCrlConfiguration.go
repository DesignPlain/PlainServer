package types

type Acmpca_getCertificateAuthorityRevocationConfigurationCrlConfiguration struct {
	// Whether the CRL is publicly readable or privately held in the CRL Amazon S3 bucket.
	S3ObjectAcl string `json:"s3ObjectAcl,omitempty" yaml:"s3ObjectAcl,omitempty"`

	// Name inserted into the certificate CRL Distribution Points extension that enables the use of an alias for the CRL distribution point.
	CustomCname string `json:"customCname,omitempty" yaml:"customCname,omitempty"`

	// Boolean value that specifies whether certificate revocation lists (CRLs) are enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Number of days until a certificate expires.
	ExpirationInDays int `json:"expirationInDays,omitempty" yaml:"expirationInDays,omitempty"`

	// Name of the S3 bucket that contains the CRL.
	S3BucketName string `json:"s3BucketName,omitempty" yaml:"s3BucketName,omitempty"`
}
