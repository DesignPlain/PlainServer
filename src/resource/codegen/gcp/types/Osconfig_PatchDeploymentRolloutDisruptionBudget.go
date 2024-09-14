package types

type Osconfig_PatchDeploymentRolloutDisruptionBudget struct {
	// Specifies a fixed value.
	Fixed int `json:"fixed,omitempty" yaml:"fixed,omitempty"`

	// Specifies the relative value defined as a percentage, which will be multiplied by a reference value.
	Percentage int `json:"percentage,omitempty" yaml:"percentage,omitempty"`
}
