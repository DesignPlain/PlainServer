package types

type S3_BucketServerSideEncryptionConfigurationRule struct {
	// A single object for setting server-side encryption by default. (documented below)
	ApplyServerSideEncryptionByDefault S3_BucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault `json:"applyServerSideEncryptionByDefault,omitempty" yaml:"applyServerSideEncryptionByDefault,omitempty"`

	// Whether or not to use [Amazon S3 Bucket Keys](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-key.html) for SSE-KMS.
	BucketKeyEnabled bool `json:"bucketKeyEnabled,omitempty" yaml:"bucketKeyEnabled,omitempty"`
}
