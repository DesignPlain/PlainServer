package types

type S3control_getMultiRegionAccessPointPublicAccessBlock struct {
	/*
	   Specifies whether Amazon S3 should restrict public bucket policies for buckets in this account. When set to `true`:
	   - Only the bucket owner and AWS Services can access buckets with public policies.
	*/
	RestrictPublicBuckets bool `json:"restrictPublicBuckets,omitempty" yaml:"restrictPublicBuckets,omitempty"`

	/*
	   Specifies whether Amazon S3 should block public access control lists (ACLs). When set to `true` causes the following behavior:
	   - PUT Bucket acl and PUT Object acl calls fail if the specified ACL is public.
	   - PUT Object calls fail if the request includes a public ACL.
	   - PUT Bucket calls fail if the request includes a public ACL.
	*/
	BlockPublicAcls bool `json:"blockPublicAcls,omitempty" yaml:"blockPublicAcls,omitempty"`

	/*
	   Specifies whether Amazon S3 should block public bucket policies for buckets in this account. When set to `true` causes Amazon S3 to:
	   - Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
	*/
	BlockPublicPolicy bool `json:"blockPublicPolicy,omitempty" yaml:"blockPublicPolicy,omitempty"`

	/*
	   Specifies whether Amazon S3 should ignore public ACLs for buckets in this account. When set to `true` causes Amazon S3 to:
	   - Ignore all public ACLs on buckets in this account and any objects that they contain.
	*/
	IgnorePublicAcls bool `json:"ignorePublicAcls,omitempty" yaml:"ignorePublicAcls,omitempty"`
}
