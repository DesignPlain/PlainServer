package types

type S3_BucketLifecycleConfigurationV2RuleAbortIncompleteMultipartUpload struct {
	// Number of days after which Amazon S3 aborts an incomplete multipart upload.
	DaysAfterInitiation int `json:"daysAfterInitiation,omitempty" yaml:"daysAfterInitiation,omitempty"`
}
