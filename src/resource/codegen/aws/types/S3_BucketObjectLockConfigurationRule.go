package types

type S3_BucketObjectLockConfigurationRule struct {
	// Default retention period that you want to apply to new objects placed in this bucket (documented below).
	DefaultRetention S3_BucketObjectLockConfigurationRuleDefaultRetention `json:"defaultRetention,omitempty" yaml:"defaultRetention,omitempty"`
}
