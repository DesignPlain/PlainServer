package types

type S3_BucketServerSideEncryptionConfiguration struct {
	// A single object for server-side encryption by default configuration. (documented below)
	Rule S3_BucketServerSideEncryptionConfigurationRule `json:"rule,omitempty" yaml:"rule,omitempty"`
}
