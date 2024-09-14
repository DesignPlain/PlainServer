package types

type Scheduler_ScheduleTargetDeadLetterConfig struct {
	// ARN of the SQS queue specified as the destination for the dead-letter queue.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`
}
