package securityhub

import types "libds/aws/types"

type AutomationRule struct {
	// An integer ranging from 1 to 1000 that represents the order in which the rule action is applied to findings. Security Hub applies rules with lower values for this parameter first.
	RuleOrder int `json:"ruleOrder,omitempty" yaml:"ruleOrder,omitempty"`

	// Whether the rule is active after it is created.
	RuleStatus string `json:"ruleStatus,omitempty" yaml:"ruleStatus,omitempty"`

	//
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A block that specifies one or more actions to update finding fields if a finding matches the conditions specified in `Criteria`. Documented below.
	Actions []types.Securityhub_AutomationRuleAction `json:"actions,omitempty" yaml:"actions,omitempty"`

	// A block that specifies a set of ASFF finding field attributes and corresponding expected values that Security Hub uses to filter findings. Documented below.
	Criteria types.Securityhub_AutomationRuleCriteria `json:"criteria,omitempty" yaml:"criteria,omitempty"`

	// The description of the rule.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Specifies whether a rule is the last to be applied with respect to a finding that matches the rule criteria. Defaults to `false`.
	IsTerminal bool `json:"isTerminal,omitempty" yaml:"isTerminal,omitempty"`

	// The name of the rule.
	RuleName string `json:"ruleName,omitempty" yaml:"ruleName,omitempty"`
}
