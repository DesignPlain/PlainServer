package types

type Kinesis_AnalyticsApplicationOutputLambda struct {
	// The ARN of the Lambda function.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`

	// The ARN of the IAM Role used to access the Lambda function.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
