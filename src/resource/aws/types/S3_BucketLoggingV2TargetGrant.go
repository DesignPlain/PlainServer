package types

type S3_BucketLoggingV2TargetGrant struct {
	// Configuration block for the person being granted permissions. See below.
	Grantee S3_BucketLoggingV2TargetGrantGrantee `json:"grantee,omitempty" yaml:"grantee,omitempty"`

	// Logging permissions assigned to the grantee for the bucket. Valid values: `FULL_CONTROL`, `READ`, `WRITE`.
	Permission string `json:"permission,omitempty" yaml:"permission,omitempty"`
}
