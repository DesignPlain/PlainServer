package types

type Lb_ListenerDefaultActionFixedResponse struct {
	// Message body.
	MessageBody string `json:"messageBody,omitempty" yaml:"messageBody,omitempty"`

	// HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.
	StatusCode string `json:"statusCode,omitempty" yaml:"statusCode,omitempty"`

	/*
	   Content type. Valid values are `text/plain`, `text/css`, `text/html`, `application/javascript` and `application/json`.

	   The following arguments are optional:
	*/
	ContentType string `json:"contentType,omitempty" yaml:"contentType,omitempty"`
}
