package types

type S3_BucketObjectLockConfigurationRuleDefaultRetention struct {
	// The number of days that you want to specify for the default retention period.
	Days int `json:"days,omitempty" yaml:"days,omitempty"`

	// The default Object Lock retention mode you want to apply to new objects placed in this bucket. Valid values are `GOVERNANCE` and `COMPLIANCE`.
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	/*
	   The number of years that you want to specify for the default retention period.

	   Either `days` or `years` must be specified, but not both.

	   > --NOTE on `object_lock_configuration`:-- You can only enable S3 Object Lock for new buckets. If you need to turn on S3 Object Lock for an existing bucket, please contact AWS Support.
	   When you create a bucket with S3 Object Lock enabled, Amazon S3 automatically enables versioning for the bucket.
	   Once you create a bucket with S3 Object Lock enabled, you can't disable Object Lock or suspend versioning for the bucket.
	*/
	Years int `json:"years,omitempty" yaml:"years,omitempty"`
}
