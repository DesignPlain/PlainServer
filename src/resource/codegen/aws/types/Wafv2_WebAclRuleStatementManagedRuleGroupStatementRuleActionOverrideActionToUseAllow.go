package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseAllow struct {
	// Defines custom handling for the web request. See `custom_request_handling` below for details.
	CustomRequestHandling Wafv2_WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseAllowCustomRequestHandling `json:"customRequestHandling,omitempty" yaml:"customRequestHandling,omitempty"`
}
