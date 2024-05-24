package types

type Cloudwatch_EventConnectionAuthParametersInvocationHttpParameters struct {
	// Contains additional body string parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
	Bodies []Cloudwatch_EventConnectionAuthParametersInvocationHttpParametersBody `json:"bodies,omitempty" yaml:"bodies,omitempty"`

	// Contains additional header parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
	Headers []Cloudwatch_EventConnectionAuthParametersInvocationHttpParametersHeader `json:"headers,omitempty" yaml:"headers,omitempty"`

	// Contains additional query string parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
	QueryStrings []Cloudwatch_EventConnectionAuthParametersInvocationHttpParametersQueryString `json:"queryStrings,omitempty" yaml:"queryStrings,omitempty"`
}
