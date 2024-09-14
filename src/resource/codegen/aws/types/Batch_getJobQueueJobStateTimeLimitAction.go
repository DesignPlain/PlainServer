package types

type Batch_getJobQueueJobStateTimeLimitAction struct {
	//
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	//
	MaxTimeSeconds int `json:"maxTimeSeconds,omitempty" yaml:"maxTimeSeconds,omitempty"`

	//
	Reason string `json:"reason,omitempty" yaml:"reason,omitempty"`

	// Describes the ability of the queue to accept new jobs (for example, `ENABLED` or `DISABLED`).
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
