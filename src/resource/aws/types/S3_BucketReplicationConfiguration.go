package types

type S3_BucketReplicationConfiguration struct {
	// The ARN of the IAM role for Amazon S3 to assume when replicating the objects.
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	// Specifies the rules managing the replication (documented below).
	Rules []S3_BucketReplicationConfigurationRule `json:"rules,omitempty" yaml:"rules,omitempty"`
}
