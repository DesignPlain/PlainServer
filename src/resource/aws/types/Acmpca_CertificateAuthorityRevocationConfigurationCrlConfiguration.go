package types

type Acmpca_CertificateAuthorityRevocationConfigurationCrlConfiguration struct {
	// Number of days until a certificate expires. Must be between 1 and 5000.
	ExpirationInDays int `json:"expirationInDays,omitempty" yaml:"expirationInDays,omitempty"`

	// Name of the S3 bucket that contains the CRL. If you do not provide a value for the `custom_cname` argument, the name of your S3 bucket is placed into the CRL Distribution Points extension of the issued certificate. You must specify a bucket policy that allows ACM PCA to write the CRL to your bucket. Must be between 3 and 255 characters in length.
	S3BucketName string `json:"s3BucketName,omitempty" yaml:"s3BucketName,omitempty"`

	// Determines whether the CRL will be publicly readable or privately held in the CRL Amazon S3 bucket. Defaults to `PUBLIC_READ`.
	S3ObjectAcl string `json:"s3ObjectAcl,omitempty" yaml:"s3ObjectAcl,omitempty"`

	// Name inserted into the certificate CRL Distribution Points extension that enables the use of an alias for the CRL distribution point. Use this value if you don't want the name of your S3 bucket to be public. Must be less than or equal to 253 characters in length.
	CustomCname string `json:"customCname,omitempty" yaml:"customCname,omitempty"`

	// Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
