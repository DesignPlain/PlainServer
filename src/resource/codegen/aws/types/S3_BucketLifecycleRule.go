package types

type S3_BucketLifecycleRule struct {
	// Specifies a period in the object's expire. See Expiration below for details.
	Expiration S3_BucketLifecycleRuleExpiration `json:"expiration,omitempty" yaml:"expiration,omitempty"`

	// Unique identifier for the rule. Must be less than or equal to 255 characters in length.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Specifies when noncurrent object versions transitions. See Noncurrent Version Transition below for details.
	NoncurrentVersionTransitions []S3_BucketLifecycleRuleNoncurrentVersionTransition `json:"noncurrentVersionTransitions,omitempty" yaml:"noncurrentVersionTransitions,omitempty"`

	// Object key prefix identifying one or more objects to which the rule applies.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
	AbortIncompleteMultipartUploadDays int `json:"abortIncompleteMultipartUploadDays,omitempty" yaml:"abortIncompleteMultipartUploadDays,omitempty"`

	// Specifies lifecycle rule status.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Specifies when noncurrent object versions expire. See Noncurrent Version Expiration below for details.
	NoncurrentVersionExpiration S3_BucketLifecycleRuleNoncurrentVersionExpiration `json:"noncurrentVersionExpiration,omitempty" yaml:"noncurrentVersionExpiration,omitempty"`

	// Specifies object tags key and value.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies a period in the object's transitions. See Transition below for details.
	Transitions []S3_BucketLifecycleRuleTransition `json:"transitions,omitempty" yaml:"transitions,omitempty"`
}
