package types

type S3_BucketObjectLockConfigurationV2Rule struct {
	// Configuration block for specifying the default Object Lock retention settings for new objects placed in the specified bucket. See below.
	DefaultRetention S3_BucketObjectLockConfigurationV2RuleDefaultRetention `json:"defaultRetention,omitempty" yaml:"defaultRetention,omitempty"`
}
