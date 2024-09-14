package types

type Iot_TopicRuleS3 struct {
	// The Amazon S3 canned ACL that controls access to the object identified by the object key. [Valid values](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#canned-acl).
	CannedAcl string `json:"cannedAcl,omitempty" yaml:"cannedAcl,omitempty"`

	// The object key.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// The ARN of the IAM role that grants access.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The Amazon S3 bucket name.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`
}
