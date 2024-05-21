package types

type Osconfig_GuestPoliciesAssignmentGroupLabel struct {
	// Google Compute Engine instance labels that must be present for an instance to be included in this assignment group.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
}
