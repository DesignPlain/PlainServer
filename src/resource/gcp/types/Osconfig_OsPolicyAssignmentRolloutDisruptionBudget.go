package types

type Osconfig_OsPolicyAssignmentRolloutDisruptionBudget struct {
	// Specifies a fixed value.
	Fixed int `json:"fixed,omitempty" yaml:"fixed,omitempty"`

	/*
	   Specifies the relative value defined as a percentage,
	   which will be multiplied by a reference value.

	   --------------------------------------------------------------------------------
	*/
	Percent int `json:"percent,omitempty" yaml:"percent,omitempty"`
}
