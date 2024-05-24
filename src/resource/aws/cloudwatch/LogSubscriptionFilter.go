package cloudwatch

type LogSubscriptionFilter struct {
	// A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events. Use empty string `""` to match everything. For more information, see the [Amazon CloudWatch Logs User Guide](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html).
	FilterPattern string `json:"filterPattern,omitempty" yaml:"filterPattern,omitempty"`

	// The name of the log group to associate the subscription filter with
	LogGroup string `json:"logGroup,omitempty" yaml:"logGroup,omitempty"`

	// A name for the subscription filter
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use `aws.lambda.Permission` resource for granting access from CloudWatch logs to the destination Lambda function.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
	DestinationArn string `json:"destinationArn,omitempty" yaml:"destinationArn,omitempty"`

	// The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are "Random" and "ByLogStream".
	Distribution string `json:"distribution,omitempty" yaml:"distribution,omitempty"`
}
