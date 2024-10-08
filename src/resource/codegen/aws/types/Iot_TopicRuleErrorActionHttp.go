package types

type Iot_TopicRuleErrorActionHttp struct {
	// Custom HTTP header IoT Core should send. It is possible to define more than one custom header.
	HttpHeaders []Iot_TopicRuleErrorActionHttpHttpHeader `json:"httpHeaders,omitempty" yaml:"httpHeaders,omitempty"`

	// The HTTPS URL.
	Url string `json:"url,omitempty" yaml:"url,omitempty"`

	// The HTTPS URL used to verify ownership of `url`.
	ConfirmationUrl string `json:"confirmationUrl,omitempty" yaml:"confirmationUrl,omitempty"`
}
