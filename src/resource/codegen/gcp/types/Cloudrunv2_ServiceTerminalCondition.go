package types

type Cloudrunv2_ServiceTerminalCondition struct {
	/*
	   (Output)
	   How to interpret failures of this condition, one of Error, Warning, Info
	*/
	Severity string `json:"severity,omitempty" yaml:"severity,omitempty"`

	/*
	   (Output)
	   State of the condition.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	/*
	   The allocation type for this traffic target.
	   Possible values are: `TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST`, `TRAFFIC_TARGET_ALLOCATION_TYPE_REVISION`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	/*
	   (Output)
	   A reason for the execution condition.
	*/
	ExecutionReason string `json:"executionReason,omitempty" yaml:"executionReason,omitempty"`

	/*
	   (Output)
	   Last time the condition transitioned from one status to another.
	   A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	*/
	LastTransitionTime string `json:"lastTransitionTime,omitempty" yaml:"lastTransitionTime,omitempty"`

	/*
	   (Output)
	   Human readable message indicating details about the current status.
	*/
	Message string `json:"message,omitempty" yaml:"message,omitempty"`

	/*
	   (Output)
	   A common (service-level) reason for this condition.
	*/
	Reason string `json:"reason,omitempty" yaml:"reason,omitempty"`

	/*
	   (Output)
	   A reason for the revision condition.
	*/
	RevisionReason string `json:"revisionReason,omitempty" yaml:"revisionReason,omitempty"`
}
