package types

type Batch_JobDefinitionRetryStrategyEvaluateOnExit struct {
	// Action to take if all of the specified conditions are met. The values are not case sensitive. Valid values: `retry`, `exit`.
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	// Glob pattern to match against the decimal representation of the exit code returned for a job.
	OnExitCode string `json:"onExitCode,omitempty" yaml:"onExitCode,omitempty"`

	// Glob pattern to match against the reason returned for a job.
	OnReason string `json:"onReason,omitempty" yaml:"onReason,omitempty"`

	// Glob pattern to match against the status reason returned for a job.
	OnStatusReason string `json:"onStatusReason,omitempty" yaml:"onStatusReason,omitempty"`
}
