package types

type Wafregional_RuleGroupActivatedRule struct {
	// Specifies the action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.
	Action Wafregional_RuleGroupActivatedRuleAction `json:"action,omitempty" yaml:"action,omitempty"`

	// Specifies the order in which the rules are evaluated. Rules with a lower value are evaluated before rules with a higher value.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// The ID of a rule
	RuleId string `json:"ruleId,omitempty" yaml:"ruleId,omitempty"`

	//
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
