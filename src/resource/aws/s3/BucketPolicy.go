package s3

type BucketPolicy struct {
	// Name of the bucket to which to apply the policy.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Text of the policy. Although this is a bucket policy rather than an IAM policy, the `aws.iam.getPolicyDocument` data source may be used, so long as it specifies a principal. For more information about building AWS IAM policy documents, see the AWS IAM Policy Document Guide. Note: Bucket policies are limited to 20 KB in size.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`
}
