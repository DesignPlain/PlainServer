package types

type Wafv2_RuleGroupRuleActionChallenge struct {
	// Defines custom handling for the web request. See Custom Request Handling below for details.
	CustomRequestHandling Wafv2_RuleGroupRuleActionChallengeCustomRequestHandling `json:"customRequestHandling,omitempty" yaml:"customRequestHandling,omitempty"`
}
