package types

type Wafv2_RuleGroupRuleActionCount struct {
	// Defines custom handling for the web request. See Custom Request Handling below for details.
	CustomRequestHandling Wafv2_RuleGroupRuleActionCountCustomRequestHandling `json:"customRequestHandling,omitempty" yaml:"customRequestHandling,omitempty"`
}
