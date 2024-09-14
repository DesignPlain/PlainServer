package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUse struct {
	//
	Allow Wafv2_WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseAllow `json:"allow,omitempty" yaml:"allow,omitempty"`

	//
	Block Wafv2_WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseBlock `json:"block,omitempty" yaml:"block,omitempty"`

	// Instructs AWS WAF to run a Captcha check against the web request. See `captcha` below for details.
	Captcha Wafv2_WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCaptcha `json:"captcha,omitempty" yaml:"captcha,omitempty"`

	// Instructs AWS WAF to run a check against the request to verify that the request is coming from a legitimate client session. See `challenge` below for details.
	Challenge Wafv2_WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseChallenge `json:"challenge,omitempty" yaml:"challenge,omitempty"`

	//
	Count Wafv2_WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCount `json:"count,omitempty" yaml:"count,omitempty"`
}
