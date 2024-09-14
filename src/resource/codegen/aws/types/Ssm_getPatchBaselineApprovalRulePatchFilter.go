package types

type Ssm_getPatchBaselineApprovalRulePatchFilter struct {
	// Key for the filter.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Value for the filter.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
