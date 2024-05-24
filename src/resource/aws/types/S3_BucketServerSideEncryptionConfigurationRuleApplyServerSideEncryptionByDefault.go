package types

type S3_BucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault struct {
	// The AWS KMS master key ID used for the SSE-KMS encryption. This can only be used when you set the value of `sse_algorithm` as `aws:kms`. The default `aws/s3` AWS KMS master key is used if this element is absent while the `sse_algorithm` is `aws:kms`.
	KmsMasterKeyId string `json:"kmsMasterKeyId,omitempty" yaml:"kmsMasterKeyId,omitempty"`

	// The server-side encryption algorithm to use. Valid values are `AES256` and `aws:kms`
	SseAlgorithm string `json:"sseAlgorithm,omitempty" yaml:"sseAlgorithm,omitempty"`
}
