package types

type Lambda_FunctionDeadLetterConfig struct {
	// ARN of an SNS topic or SQS queue to notify when an invocation fails. If this option is used, the function's IAM role must be granted suitable access to write to the target object, which means allowing either the `sns:Publish` or `sqs:SendMessage` action on this ARN, depending on which service is targeted.
	TargetArn string `json:"targetArn,omitempty" yaml:"targetArn,omitempty"`
}
