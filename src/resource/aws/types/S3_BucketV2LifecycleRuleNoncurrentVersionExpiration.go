package types

type S3_BucketV2LifecycleRuleNoncurrentVersionExpiration struct {
	// Specifies the number of days noncurrent object versions expire.
	Days int `json:"days,omitempty" yaml:"days,omitempty"`
}
