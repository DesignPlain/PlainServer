package types

type Wafv2_WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseAllowCustomRequestHandlingInsertHeader struct {
	// Name of the custom header. For custom request header insertion, when AWS WAF inserts the header into the request, it prefixes this name `x-amzn-waf-`, to avoid confusion with the headers that are already in the request. For example, for the header name `sample`, AWS WAF inserts the header `x-amzn-waf-sample`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Value of the custom header.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
