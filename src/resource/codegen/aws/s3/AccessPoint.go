package s3

import types "libds/aws/types"

type AccessPoint struct {
	// Configuration block to manage the `PublicAccessBlock` configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. Detailed below.
	PublicAccessBlockConfiguration types.S3_AccessPointPublicAccessBlockConfiguration `json:"publicAccessBlockConfiguration,omitempty" yaml:"publicAccessBlockConfiguration,omitempty"`

	// Configuration block to restrict access to this access point to requests from the specified Virtual Private Cloud (VPC). Required for S3 on Outposts. Detailed below.
	VpcConfiguration types.S3_AccessPointVpcConfiguration `json:"vpcConfiguration,omitempty" yaml:"vpcConfiguration,omitempty"`

	// AWS account ID for the owner of the bucket for which you want to create an access point. Defaults to automatically determined account ID of the AWS provider.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// Name of an AWS Partition S3 General Purpose Bucket or the ARN of S3 on Outposts Bucket that you want to associate this access point with.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// AWS account ID associated with the S3 bucket associated with this access point.
	BucketAccountId string `json:"bucketAccountId,omitempty" yaml:"bucketAccountId,omitempty"`

	/*
	   Name you want to assign to this access point. See the [AWS documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/creating-access-points.html?icmpid=docs_amazons3_console#access-points-names) for naming conditions.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Valid JSON document that specifies the policy that you want to apply to this access point. Removing `policy` from your configuration or setting `policy` to null or an empty string (i.e., `policy = ""`) _will not_ delete the policy since it could have been set by `aws.s3control.AccessPointPolicy`. To remove the `policy`, set it to `"{}"` (an empty JSON document).
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`
}
