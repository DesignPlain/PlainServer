package types

type Pipes_PipeSourceParametersDynamodbStreamParametersDeadLetterConfig struct {
	// The ARN of the Amazon SQS queue specified as the target for the dead-letter queue.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`
}
