package types

type S3_BucketV2ServerSideEncryptionConfiguration struct {
	// Single object for server-side encryption by default configuration. (documented below)
	Rules []S3_BucketV2ServerSideEncryptionConfigurationRule `json:"rules,omitempty" yaml:"rules,omitempty"`
}
