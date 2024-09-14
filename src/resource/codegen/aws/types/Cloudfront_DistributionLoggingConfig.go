package types

type Cloudfront_DistributionLoggingConfig struct {
	// Amazon S3 bucket to store the access logs in, for example, `myawslogbucket.s3.amazonaws.com`.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Whether to include cookies in access logs (default: `false`).
	IncludeCookies bool `json:"includeCookies,omitempty" yaml:"includeCookies,omitempty"`

	// Prefix to the access log filenames for this distribution, for example, `myprefix/`.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
}
