package types

type Iot_TopicRuleS3 struct {
	// The Amazon S3 bucket name.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// The Amazon S3 canned ACL that controls access to the object identified by the object key. [Valid values](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#canned-acl).
	CannedAcl string `json:"cannedAcl,omitempty" yaml:"cannedAcl,omitempty"`

	// The name of the HTTP header.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// The IAM role ARN that allows access to the CloudWatch alarm.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
