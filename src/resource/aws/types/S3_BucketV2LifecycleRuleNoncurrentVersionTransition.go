package types

type S3_BucketV2LifecycleRuleNoncurrentVersionTransition struct {
	// Specifies the Amazon S3 [storage class](https://docs.aws.amazon.com/AmazonS3/latest/API/API_Transition.html#AmazonS3-Type-Transition-StorageClass) to which you want the object to transition.
	StorageClass string `json:"storageClass,omitempty" yaml:"storageClass,omitempty"`

	// Specifies the number of days noncurrent object versions transition.
	Days int `json:"days,omitempty" yaml:"days,omitempty"`
}
