package types

type Networkfirewall_RuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsCustomAction struct {
	// A configuration block describing the custom action associated with the `action_name`. See Action Definition below for details.
	ActionDefinition Networkfirewall_RuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsCustomActionActionDefinition `json:"actionDefinition,omitempty" yaml:"actionDefinition,omitempty"`

	// A friendly name of the custom action.
	ActionName string `json:"actionName,omitempty" yaml:"actionName,omitempty"`
}
