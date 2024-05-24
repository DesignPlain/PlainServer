package types

type Wafv2_WebAclRuleActionBlock struct {
	// Defines a custom response for the web request. See `custom_response` below for details.
	CustomResponse Wafv2_WebAclRuleActionBlockCustomResponse `json:"customResponse,omitempty" yaml:"customResponse,omitempty"`
}
