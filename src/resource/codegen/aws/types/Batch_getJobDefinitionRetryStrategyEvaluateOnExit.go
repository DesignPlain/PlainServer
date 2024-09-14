package types

type Batch_getJobDefinitionRetryStrategyEvaluateOnExit struct {
	// Contains a glob pattern to match against the Reason returned for a job.
	OnReason string `json:"onReason,omitempty" yaml:"onReason,omitempty"`

	// Contains a glob pattern to match against the StatusReason returned for a job.
	OnStatusReason string `json:"onStatusReason,omitempty" yaml:"onStatusReason,omitempty"`

	// Specifies the action to take if all of the specified conditions (onStatusReason, onReason, and onExitCode) are met. The values aren't case sensitive.
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	// Contains a glob pattern to match against the decimal representation of the ExitCode returned for a job.
	OnExitCode string `json:"onExitCode,omitempty" yaml:"onExitCode,omitempty"`
}
