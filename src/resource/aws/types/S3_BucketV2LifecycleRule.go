package types

type S3_BucketV2LifecycleRule struct {
	// Specifies when noncurrent object versions expire. See Noncurrent Version Expiration below for details.
	NoncurrentVersionExpirations []S3_BucketV2LifecycleRuleNoncurrentVersionExpiration `json:"noncurrentVersionExpirations,omitempty" yaml:"noncurrentVersionExpirations,omitempty"`

	// Specifies when noncurrent object versions transitions. See Noncurrent Version Transition below for details.
	NoncurrentVersionTransitions []S3_BucketV2LifecycleRuleNoncurrentVersionTransition `json:"noncurrentVersionTransitions,omitempty" yaml:"noncurrentVersionTransitions,omitempty"`

	// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
	AbortIncompleteMultipartUploadDays int `json:"abortIncompleteMultipartUploadDays,omitempty" yaml:"abortIncompleteMultipartUploadDays,omitempty"`

	// Specifies a period in the object's expire. See Expiration below for details.
	Expirations []S3_BucketV2LifecycleRuleExpiration `json:"expirations,omitempty" yaml:"expirations,omitempty"`

	// Object key prefix identifying one or more objects to which the rule applies.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// Specifies object tags key and value.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies a period in the object's transitions. See Transition below for details.
	Transitions []S3_BucketV2LifecycleRuleTransition `json:"transitions,omitempty" yaml:"transitions,omitempty"`

	// Specifies lifecycle rule status.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Unique identifier for the rule. Must be less than or equal to 255 characters in length.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`
}
