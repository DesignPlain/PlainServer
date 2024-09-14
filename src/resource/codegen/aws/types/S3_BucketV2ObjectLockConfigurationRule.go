package types

type S3_BucketV2ObjectLockConfigurationRule struct {
	// Default retention period that you want to apply to new objects placed in this bucket (documented below).
	DefaultRetentions []S3_BucketV2ObjectLockConfigurationRuleDefaultRetention `json:"defaultRetentions,omitempty" yaml:"defaultRetentions,omitempty"`
}
