package types

type Cloudformation_getCloudFormationTypeLoggingConfig struct {
	// Name of the CloudWatch Log Group where CloudFormation sends error logging information when invoking the type's handlers.
	LogGroupName string `json:"logGroupName,omitempty" yaml:"logGroupName,omitempty"`

	// ARN of the IAM Role CloudFormation assumes when sending error logging information to CloudWatch Logs.
	LogRoleArn string `json:"logRoleArn,omitempty" yaml:"logRoleArn,omitempty"`
}
