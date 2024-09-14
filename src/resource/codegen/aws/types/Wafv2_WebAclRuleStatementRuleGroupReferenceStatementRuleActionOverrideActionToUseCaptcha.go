package types

type Wafv2_WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCaptcha struct {
	// Defines custom handling for the web request. See `custom_request_handling` below for details.
	CustomRequestHandling Wafv2_WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCaptchaCustomRequestHandling `json:"customRequestHandling,omitempty" yaml:"customRequestHandling,omitempty"`
}
