package sns

type TopicPolicy struct {
	// The ARN of the SNS topic
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// The fully-formed AWS policy as JSON.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`
}
