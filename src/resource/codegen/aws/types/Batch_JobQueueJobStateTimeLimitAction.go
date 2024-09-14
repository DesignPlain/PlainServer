package types

type Batch_JobQueueJobStateTimeLimitAction struct {
	/*
	   The action to take when a job is at the head of the job queue in the specified state for the specified period of time. Valid values include `"CANCEL"`
	   - `job_state_time_limit_action.#.max_time_seconds` - The approximate amount of time, in seconds, that must pass with the job in the specified state before the action is taken. Valid values include integers between `600` & `86400`
	*/
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	//
	MaxTimeSeconds int `json:"maxTimeSeconds,omitempty" yaml:"maxTimeSeconds,omitempty"`

	// The reason to log for the action being taken.
	Reason string `json:"reason,omitempty" yaml:"reason,omitempty"`

	// The state of the job needed to trigger the action. Valid values include `"RUNNABLE"`.
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
