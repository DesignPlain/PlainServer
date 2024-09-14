package types

type Batch_JobDefinitionTimeout struct {
	// Time duration in seconds after which AWS Batch terminates your jobs if they have not finished. The minimum value for the timeout is `60` seconds.
	AttemptDurationSeconds int `json:"attemptDurationSeconds,omitempty" yaml:"attemptDurationSeconds,omitempty"`
}
