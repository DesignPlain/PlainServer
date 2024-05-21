package types

type Clouddeploy_AutomationRule struct {
	/*
	   Optional. The `AdvanceRolloutRule` will automatically advance a successful Rollout.
	   Structure is documented below.
	*/
	AdvanceRolloutRule Clouddeploy_AutomationRuleAdvanceRolloutRule `json:"advanceRolloutRule,omitempty" yaml:"advanceRolloutRule,omitempty"`

	/*
	   Optional. `PromoteReleaseRule` will automatically promote a release from the current target to a specified target.
	   Structure is documented below.
	*/
	PromoteReleaseRule Clouddeploy_AutomationRulePromoteReleaseRule `json:"promoteReleaseRule,omitempty" yaml:"promoteReleaseRule,omitempty"`
}
