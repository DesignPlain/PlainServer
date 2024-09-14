package types

type Scheduler_ScheduleTargetKinesisParameters struct {
	// Specifies the shard to which EventBridge Scheduler sends the event. Up to 256 characters.
	PartitionKey string `json:"partitionKey,omitempty" yaml:"partitionKey,omitempty"`
}
