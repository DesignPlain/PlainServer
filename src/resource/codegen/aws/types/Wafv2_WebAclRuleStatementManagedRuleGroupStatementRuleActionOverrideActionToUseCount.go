package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCount struct {
	// Defines custom handling for the web request. See `custom_request_handling` below for details.
	CustomRequestHandling Wafv2_WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandling `json:"customRequestHandling,omitempty" yaml:"customRequestHandling,omitempty"`
}
