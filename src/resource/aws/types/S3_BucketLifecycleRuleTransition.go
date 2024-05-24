package types

type S3_BucketLifecycleRuleTransition struct {
	// Specifies the date after which you want the corresponding action to take effect.
	Date string `json:"date,omitempty" yaml:"date,omitempty"`

	// Specifies the number of days after object creation when the specific rule action takes effect.
	Days int `json:"days,omitempty" yaml:"days,omitempty"`

	// Specifies the Amazon S3 [storage class](https://docs.aws.amazon.com/AmazonS3/latest/API/API_Transition.html#AmazonS3-Type-Transition-StorageClass) to which you want the object to transition.
	StorageClass string `json:"storageClass,omitempty" yaml:"storageClass,omitempty"`
}
