package types

type Wafv2_RuleGroupRuleActionAllow struct {
	// Defines custom handling for the web request. See Custom Request Handling below for details.
	CustomRequestHandling Wafv2_RuleGroupRuleActionAllowCustomRequestHandling `json:"customRequestHandling,omitempty" yaml:"customRequestHandling,omitempty"`
}
