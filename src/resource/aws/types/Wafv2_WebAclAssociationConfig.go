package types

type Wafv2_WebAclAssociationConfig struct {
	// Customizes the request body that your protected resource forward to AWS WAF for inspection. See `request_body` below for details.
	RequestBodies []Wafv2_WebAclAssociationConfigRequestBody `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`
}
