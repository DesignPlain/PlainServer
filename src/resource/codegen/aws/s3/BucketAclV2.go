package s3

import types "libds/aws/types"

type BucketAclV2 struct {
	// Configuration block that sets the ACL permissions for an object per grantee. See below.
	AccessControlPolicy types.S3_BucketAclV2AccessControlPolicy `json:"accessControlPolicy,omitempty" yaml:"accessControlPolicy,omitempty"`

	// Canned ACL to apply to the bucket.
	Acl string `json:"acl,omitempty" yaml:"acl,omitempty"`

	// Bucket to which to apply the ACL.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Account ID of the expected bucket owner.
	ExpectedBucketOwner string `json:"expectedBucketOwner,omitempty" yaml:"expectedBucketOwner,omitempty"`
}
