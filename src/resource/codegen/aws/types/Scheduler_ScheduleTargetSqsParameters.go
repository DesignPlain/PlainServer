package types

type Scheduler_ScheduleTargetSqsParameters struct {
	// FIFO message group ID to use as the target.
	MessageGroupId string `json:"messageGroupId,omitempty" yaml:"messageGroupId,omitempty"`
}
