package types

type Acmpca_getCertificateAuthorityRevocationConfigurationCrlConfiguration struct {
	//
	S3BucketName string `json:"s3BucketName,omitempty" yaml:"s3BucketName,omitempty"`

	//
	S3ObjectAcl string `json:"s3ObjectAcl,omitempty" yaml:"s3ObjectAcl,omitempty"`

	//
	CustomCname string `json:"customCname,omitempty" yaml:"customCname,omitempty"`

	//
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	//
	ExpirationInDays int `json:"expirationInDays,omitempty" yaml:"expirationInDays,omitempty"`
}
