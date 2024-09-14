package types

type Cloudrun_getServiceStatusCondition struct {
	// Human readable message indicating details about the current status.
	Message string `json:"message,omitempty" yaml:"message,omitempty"`

	// One-word CamelCase reason for the condition's current status.
	Reason string `json:"reason,omitempty" yaml:"reason,omitempty"`

	// Status of the condition, one of True, False, Unknown.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// Type of domain mapping condition.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
