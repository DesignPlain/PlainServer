package types

type Osconfig_OsPolicyAssignmentRollout struct {
	/*
	   The maximum number (or percentage) of VMs
	   per zone to disrupt at any given moment. Structure is
	   documented below.
	*/
	DisruptionBudget Osconfig_OsPolicyAssignmentRolloutDisruptionBudget `json:"disruptionBudget,omitempty" yaml:"disruptionBudget,omitempty"`

	/*
	   This determines the minimum duration of
	   time to wait after the configuration changes are applied through the current
	   rollout. A VM continues to count towards the `disruption_budget` at least
	   until this duration of time has passed after configuration changes are
	   applied.
	*/
	MinWaitDuration string `json:"minWaitDuration,omitempty" yaml:"minWaitDuration,omitempty"`
}
