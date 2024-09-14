package types

type Ses_ReceiptRuleS3Action struct {
	// The name of the S3 bucket
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// The ARN of the KMS key
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	// The key prefix of the S3 bucket
	ObjectKeyPrefix string `json:"objectKeyPrefix,omitempty" yaml:"objectKeyPrefix,omitempty"`

	// The position of the action in the receipt rule
	Position int `json:"position,omitempty" yaml:"position,omitempty"`

	// The ARN of an SNS topic to notify
	TopicArn string `json:"topicArn,omitempty" yaml:"topicArn,omitempty"`
}
