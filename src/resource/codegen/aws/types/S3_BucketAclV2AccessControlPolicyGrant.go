package types

type S3_BucketAclV2AccessControlPolicyGrant struct {
	// Configuration block for the person being granted permissions. See below.
	Grantee S3_BucketAclV2AccessControlPolicyGrantGrantee `json:"grantee,omitempty" yaml:"grantee,omitempty"`

	// Logging permissions assigned to the grantee for the bucket. Valid values: `FULL_CONTROL`, `WRITE`, `WRITE_ACP`, `READ`, `READ_ACP`. See [What permissions can I grant?](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#permissions) for more details about what each permission means in the context of buckets.
	Permission string `json:"permission,omitempty" yaml:"permission,omitempty"`
}
