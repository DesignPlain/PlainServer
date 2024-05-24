package types

type Cloudwatch_EventTargetDeadLetterConfig struct {
	// ARN of the SQS queue specified as the target for the dead-letter queue.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`
}
