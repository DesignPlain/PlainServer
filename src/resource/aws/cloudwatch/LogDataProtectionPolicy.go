package cloudwatch

type LogDataProtectionPolicy struct {
	// The name of the log group under which the log stream is to be created.
	LogGroupName string `json:"logGroupName,omitempty" yaml:"logGroupName,omitempty"`

	// Specifies the data protection policy in JSON. Read more at [Data protection policy syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data-start.html#mask-sensitive-log-data-policysyntax).
	PolicyDocument string `json:"policyDocument,omitempty" yaml:"policyDocument,omitempty"`
}
