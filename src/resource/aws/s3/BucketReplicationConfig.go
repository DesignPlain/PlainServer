package s3

import types "DesignSphere_Server/src/resource/aws/types"

type BucketReplicationConfig struct {
	// Name of the source S3 bucket you want Amazon S3 to monitor.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// ARN of the IAM role for Amazon S3 to assume when replicating the objects.
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	// List of configuration blocks describing the rules managing the replication. See below.
	Rules []types.S3_BucketReplicationConfigRule `json:"rules,omitempty" yaml:"rules,omitempty"`

	/*
	   Token to allow replication to be enabled on an Object Lock-enabled bucket. You must contact AWS support for the bucket's "Object Lock token".
	   For more details, see [Using S3 Object Lock with replication](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-managing.html#object-lock-managing-replication).
	*/
	Token string `json:"token,omitempty" yaml:"token,omitempty"`
}
