package types

type Cloudrunv2_getServiceTerminalCondition struct {
	// How to interpret failures of this condition, one of Error, Warning, Info
	Severity string `json:"severity,omitempty" yaml:"severity,omitempty"`

	// State of the condition.
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	// type is used to communicate the status of the reconciliation process. See also: https://github.com/knative/serving/blob/main/docs/spec/errors.md#error-conditions-and-reporting Types common to all resources include: - "Ready": True when the Resource is ready.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// A reason for the execution condition.
	ExecutionReason string `json:"executionReason,omitempty" yaml:"executionReason,omitempty"`

	// Last time the condition transitioned from one status to another.
	LastTransitionTime string `json:"lastTransitionTime,omitempty" yaml:"lastTransitionTime,omitempty"`

	// Human readable message indicating details about the current status.
	Message string `json:"message,omitempty" yaml:"message,omitempty"`

	// A common (service-level) reason for this condition.
	Reason string `json:"reason,omitempty" yaml:"reason,omitempty"`

	// A reason for the revision condition.
	RevisionReason string `json:"revisionReason,omitempty" yaml:"revisionReason,omitempty"`
}
