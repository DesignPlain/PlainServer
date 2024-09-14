package types

type S3control_BucketLifecycleConfigurationRule struct {
	// Configuration block containing settings for expiration of objects.
	Expiration S3control_BucketLifecycleConfigurationRuleExpiration `json:"expiration,omitempty" yaml:"expiration,omitempty"`

	// Configuration block containing settings for filtering.
	Filter S3control_BucketLifecycleConfigurationRuleFilter `json:"filter,omitempty" yaml:"filter,omitempty"`

	// Unique identifier for the rule.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Status of the rule. Valid values: `Enabled` and `Disabled`. Defaults to `Enabled`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// Configuration block containing settings for abort incomplete multipart upload.
	AbortIncompleteMultipartUpload S3control_BucketLifecycleConfigurationRuleAbortIncompleteMultipartUpload `json:"abortIncompleteMultipartUpload,omitempty" yaml:"abortIncompleteMultipartUpload,omitempty"`
}
