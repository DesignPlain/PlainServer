package types

type Wafv2_WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUse struct {
	// Instructs AWS WAF to run a check against the request to verify that the request is coming from a legitimate client session. See `challenge` below for details.
	Challenge Wafv2_WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallenge `json:"challenge,omitempty" yaml:"challenge,omitempty"`

	//
	Count Wafv2_WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCount `json:"count,omitempty" yaml:"count,omitempty"`

	//
	Allow Wafv2_WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseAllow `json:"allow,omitempty" yaml:"allow,omitempty"`

	//
	Block Wafv2_WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseBlock `json:"block,omitempty" yaml:"block,omitempty"`

	// Instructs AWS WAF to run a Captcha check against the web request. See `captcha` below for details.
	Captcha Wafv2_WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseCaptcha `json:"captcha,omitempty" yaml:"captcha,omitempty"`
}
