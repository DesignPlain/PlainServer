package types

type S3_BucketV2ServerSideEncryptionConfigurationRule struct {
	// Single object for setting server-side encryption by default. (documented below)
	ApplyServerSideEncryptionByDefaults []S3_BucketV2ServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault `json:"applyServerSideEncryptionByDefaults,omitempty" yaml:"applyServerSideEncryptionByDefaults,omitempty"`

	// Whether or not to use [Amazon S3 Bucket Keys](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-key.html) for SSE-KMS.
	BucketKeyEnabled bool `json:"bucketKeyEnabled,omitempty" yaml:"bucketKeyEnabled,omitempty"`
}
