package types

type Wafv2_WebAclCustomResponseBody struct {
	// Payload of the custom response.
	Content string `json:"content,omitempty" yaml:"content,omitempty"`

	// Type of content in the payload that you are defining in the `content` argument. Valid values are `TEXT_PLAIN`, `TEXT_HTML`, or `APPLICATION_JSON`.
	ContentType string `json:"contentType,omitempty" yaml:"contentType,omitempty"`

	// Unique key identifying the custom response body. This is referenced by the `custom_response_body_key` argument in the `custom_response` block.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
