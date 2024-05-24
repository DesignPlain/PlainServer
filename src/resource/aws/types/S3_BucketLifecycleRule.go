package types

type S3_BucketLifecycleRule struct {
	// Specifies when noncurrent object versions expire (documented below).
	NoncurrentVersionExpiration S3_BucketLifecycleRuleNoncurrentVersionExpiration `json:"noncurrentVersionExpiration,omitempty" yaml:"noncurrentVersionExpiration,omitempty"`

	// Object key prefix identifying one or more objects to which the rule applies.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// Specifies object tags key and value.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies a period in the object's transitions (documented below).
	Transitions []S3_BucketLifecycleRuleTransition `json:"transitions,omitempty" yaml:"transitions,omitempty"`

	// Specifies lifecycle rule status.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Specifies a period in the object's expire (documented below).
	Expiration S3_BucketLifecycleRuleExpiration `json:"expiration,omitempty" yaml:"expiration,omitempty"`

	/*
	   Specifies when noncurrent object versions transitions (documented below).

	   At least one of `abort_incomplete_multipart_upload_days`, `expiration`, `transition`, `noncurrent_version_expiration`, `noncurrent_version_transition` must be specified.
	*/
	NoncurrentVersionTransitions []S3_BucketLifecycleRuleNoncurrentVersionTransition `json:"noncurrentVersionTransitions,omitempty" yaml:"noncurrentVersionTransitions,omitempty"`

	// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
	AbortIncompleteMultipartUploadDays int `json:"abortIncompleteMultipartUploadDays,omitempty" yaml:"abortIncompleteMultipartUploadDays,omitempty"`

	// Unique identifier for the rule. Must be less than or equal to 255 characters in length.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`
}
