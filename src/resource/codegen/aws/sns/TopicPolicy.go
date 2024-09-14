package sns

type TopicPolicy struct {
	// The fully-formed AWS policy as JSON.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// The ARN of the SNS topic
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`
}
