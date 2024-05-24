package types

type Wafv2_WebAclDefaultActionBlock struct {
	// Defines a custom response for the web request. See `custom_response` below for details.
	CustomResponse Wafv2_WebAclDefaultActionBlockCustomResponse `json:"customResponse,omitempty" yaml:"customResponse,omitempty"`
}
