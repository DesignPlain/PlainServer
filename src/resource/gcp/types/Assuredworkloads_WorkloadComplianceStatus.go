package types

type Assuredworkloads_WorkloadComplianceStatus struct {
	// Number of current orgPolicy violations which are acknowledged.
	AcknowledgedViolationCounts []int `json:"acknowledgedViolationCounts,omitempty" yaml:"acknowledgedViolationCounts,omitempty"`

	// Number of current orgPolicy violations which are not acknowledged.
	ActiveViolationCounts []int `json:"activeViolationCounts,omitempty" yaml:"activeViolationCounts,omitempty"`
}
