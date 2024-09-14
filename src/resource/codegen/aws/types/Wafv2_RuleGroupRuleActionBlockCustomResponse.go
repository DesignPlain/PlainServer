package types

type Wafv2_RuleGroupRuleActionBlockCustomResponse struct {
	// References the response body that you want AWS WAF to return to the web request client. This must reference a `key` defined in a `custom_response_body` block of this resource.
	CustomResponseBodyKey string `json:"customResponseBodyKey,omitempty" yaml:"customResponseBodyKey,omitempty"`

	// The HTTP status code to return to the client.
	ResponseCode int `json:"responseCode,omitempty" yaml:"responseCode,omitempty"`

	// The `response_header` blocks used to define the HTTP response headers added to the response. See Custom HTTP Header below for details.
	ResponseHeaders []Wafv2_RuleGroupRuleActionBlockCustomResponseResponseHeader `json:"responseHeaders,omitempty" yaml:"responseHeaders,omitempty"`
}
