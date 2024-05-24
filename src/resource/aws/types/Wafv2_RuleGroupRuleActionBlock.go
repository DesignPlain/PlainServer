package types

type Wafv2_RuleGroupRuleActionBlock struct {
	// Defines a custom response for the web request. See Custom Response below for details.
	CustomResponse Wafv2_RuleGroupRuleActionBlockCustomResponse `json:"customResponse,omitempty" yaml:"customResponse,omitempty"`
}
