package types

type Batch_getJobDefinitionTimeout struct {
	// The job timeout time (in seconds) that's measured from the job attempt's startedAt timestamp.
	AttemptDurationSeconds int `json:"attemptDurationSeconds,omitempty" yaml:"attemptDurationSeconds,omitempty"`
}
