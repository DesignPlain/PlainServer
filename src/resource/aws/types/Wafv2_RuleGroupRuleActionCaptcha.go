package types

type Wafv2_RuleGroupRuleActionCaptcha struct {
	// Defines custom handling for the web request. See Custom Request Handling below for details.
	CustomRequestHandling Wafv2_RuleGroupRuleActionCaptchaCustomRequestHandling `json:"customRequestHandling,omitempty" yaml:"customRequestHandling,omitempty"`
}
