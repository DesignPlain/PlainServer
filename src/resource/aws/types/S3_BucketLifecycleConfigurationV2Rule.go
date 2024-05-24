package types

type S3_BucketLifecycleConfigurationV2Rule struct {
	// Configuration block that specifies the days since the initiation of an incomplete multipart upload that Amazon S3 will wait before permanently removing all parts of the upload. See below.
	AbortIncompleteMultipartUpload S3_BucketLifecycleConfigurationV2RuleAbortIncompleteMultipartUpload `json:"abortIncompleteMultipartUpload,omitempty" yaml:"abortIncompleteMultipartUpload,omitempty"`

	// Set of configuration blocks that specify when an Amazon S3 object transitions to a specified storage class. See below.
	Transitions []S3_BucketLifecycleConfigurationV2RuleTransition `json:"transitions,omitempty" yaml:"transitions,omitempty"`

	// Set of configuration blocks that specify the transition rule for the lifecycle rule that describes when noncurrent objects transition to a specific storage class. See below.
	NoncurrentVersionTransitions []S3_BucketLifecycleConfigurationV2RuleNoncurrentVersionTransition `json:"noncurrentVersionTransitions,omitempty" yaml:"noncurrentVersionTransitions,omitempty"`

	// --DEPRECATED-- Use `filter` instead. This has been deprecated by Amazon S3. Prefix identifying one or more objects to which the rule applies. Defaults to an empty string (`""`) if `filter` is not specified.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// Whether the rule is currently being applied. Valid values: `Enabled` or `Disabled`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// Configuration block that specifies the expiration for the lifecycle of the object in the form of date, days and, whether the object has a delete marker. See below.
	Expiration S3_BucketLifecycleConfigurationV2RuleExpiration `json:"expiration,omitempty" yaml:"expiration,omitempty"`

	// Configuration block used to identify objects that a Lifecycle Rule applies to. See below. If not specified, the `rule` will default to using `prefix`.
	Filter S3_BucketLifecycleConfigurationV2RuleFilter `json:"filter,omitempty" yaml:"filter,omitempty"`

	// Unique identifier for the rule. The value cannot be longer than 255 characters.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Configuration block that specifies when noncurrent object versions expire. See below.
	NoncurrentVersionExpiration S3_BucketLifecycleConfigurationV2RuleNoncurrentVersionExpiration `json:"noncurrentVersionExpiration,omitempty" yaml:"noncurrentVersionExpiration,omitempty"`
}
