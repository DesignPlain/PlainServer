package types

type S3_BucketObjectLockConfigurationRule struct {
	// The default retention period that you want to apply to new objects placed in this bucket.
	DefaultRetention S3_BucketObjectLockConfigurationRuleDefaultRetention `json:"defaultRetention,omitempty" yaml:"defaultRetention,omitempty"`
}
