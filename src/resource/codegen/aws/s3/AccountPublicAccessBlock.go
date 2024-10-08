package s3

type AccountPublicAccessBlock struct {
	/*
	   Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
	   - Ignore all public ACLs on buckets in this account and any objects that they contain.
	*/
	IgnorePublicAcls bool `json:"ignorePublicAcls,omitempty" yaml:"ignorePublicAcls,omitempty"`

	/*
	   Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
	   - Only the bucket owner and AWS Services can access buckets with public policies.
	*/
	RestrictPublicBuckets bool `json:"restrictPublicBuckets,omitempty" yaml:"restrictPublicBuckets,omitempty"`

	// AWS account ID to configure. Defaults to automatically determined account ID of the this provider AWS provider.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	/*
	   Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
	   - PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
	   - PUT Object calls fail if the request includes a public ACL.
	*/
	BlockPublicAcls bool `json:"blockPublicAcls,omitempty" yaml:"blockPublicAcls,omitempty"`

	/*
	   Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing bucket policies. When set to `true` causes Amazon S3 to:
	   - Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
	*/
	BlockPublicPolicy bool `json:"blockPublicPolicy,omitempty" yaml:"blockPublicPolicy,omitempty"`
}
