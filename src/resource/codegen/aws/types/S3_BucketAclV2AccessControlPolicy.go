package types

type S3_BucketAclV2AccessControlPolicy struct {
	// Set of `grant` configuration blocks. See below.
	Grants []S3_BucketAclV2AccessControlPolicyGrant `json:"grants,omitempty" yaml:"grants,omitempty"`

	// Configuration block for the bucket owner's display name and ID. See below.
	Owner S3_BucketAclV2AccessControlPolicyOwner `json:"owner,omitempty" yaml:"owner,omitempty"`
}
