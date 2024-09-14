package cloudwatch

type LogDestination struct {
	// A name for the log destination.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to put data into the target.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The ARN of the target Amazon Kinesis stream resource for the destination.
	TargetArn string `json:"targetArn,omitempty" yaml:"targetArn,omitempty"`
}
