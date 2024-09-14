package types

type S3_BucketObjectLockConfigurationRuleDefaultRetention struct {
	// Number of days that you want to specify for the default retention period.
	Days int `json:"days,omitempty" yaml:"days,omitempty"`

	// Default Object Lock retention mode you want to apply to new objects placed in this bucket. Valid values are `GOVERNANCE` and `COMPLIANCE`.
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	// Number of years that you want to specify for the default retention period.
	Years int `json:"years,omitempty" yaml:"years,omitempty"`
}
