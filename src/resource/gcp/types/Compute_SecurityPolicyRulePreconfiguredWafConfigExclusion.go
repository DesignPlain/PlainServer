package types

type Compute_SecurityPolicyRulePreconfiguredWafConfigExclusion struct {
	// Request header whose value will be excluded from inspection during preconfigured WAF evaluation. Structure is documented below.
	RequestHeaders []Compute_SecurityPolicyRulePreconfiguredWafConfigExclusionRequestHeader `json:"requestHeaders,omitempty" yaml:"requestHeaders,omitempty"`

	// Request URI from the request line to be excluded from inspection during preconfigured WAF evaluation. When specifying this field, the query or fragment part should be excluded. Structure is documented below.
	RequestQueryParams []Compute_SecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParam `json:"requestQueryParams,omitempty" yaml:"requestQueryParams,omitempty"`

	// Request query parameter whose value will be excluded from inspection during preconfigured WAF evaluation. Note that the parameter can be in the query string or in the POST body. Structure is documented below.
	RequestUris []Compute_SecurityPolicyRulePreconfiguredWafConfigExclusionRequestUri `json:"requestUris,omitempty" yaml:"requestUris,omitempty"`

	/*
	   A list of target rule IDs under the WAF rule set to apply the preconfigured WAF exclusion. If omitted, it refers to all the rule IDs under the WAF rule set.

	   <a name="nested_field_params"></a>The `request_header`, `request_cookie`, `request_uri` and `request_query_param` blocks support:
	*/
	TargetRuleIds []string `json:"targetRuleIds,omitempty" yaml:"targetRuleIds,omitempty"`

	// Target WAF rule set to apply the preconfigured WAF exclusion.
	TargetRuleSet string `json:"targetRuleSet,omitempty" yaml:"targetRuleSet,omitempty"`

	// Request cookie whose value will be excluded from inspection during preconfigured WAF evaluation. Structure is documented below.
	RequestCookies []Compute_SecurityPolicyRulePreconfiguredWafConfigExclusionRequestCooky `json:"requestCookies,omitempty" yaml:"requestCookies,omitempty"`
}
