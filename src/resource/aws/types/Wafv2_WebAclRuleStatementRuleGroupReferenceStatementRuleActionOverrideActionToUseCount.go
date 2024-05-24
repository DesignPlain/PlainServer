package types

type Wafv2_WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCount struct {
	// Defines custom handling for the web request. See `custom_request_handling` below for details.
	CustomRequestHandling Wafv2_WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCountCustomRequestHandling `json:"customRequestHandling,omitempty" yaml:"customRequestHandling,omitempty"`
}
