package types

type Lb_ListenerRuleActionFixedResponse struct {
	// The content type. Valid values are `text/plain`, `text/css`, `text/html`, `application/javascript` and `application/json`.
	ContentType string `json:"contentType,omitempty" yaml:"contentType,omitempty"`

	// The message body.
	MessageBody string `json:"messageBody,omitempty" yaml:"messageBody,omitempty"`

	// The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.
	StatusCode string `json:"statusCode,omitempty" yaml:"statusCode,omitempty"`
}
