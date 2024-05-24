package types

type Pipes_PipeEnrichmentParametersHttpParameters struct {
	// Key-value mapping of the headers that need to be sent as part of request invoking the API Gateway REST API or EventBridge ApiDestination.
	HeaderParameters map[string]string `json:"headerParameters,omitempty" yaml:"headerParameters,omitempty"`

	// The path parameter values to be used to populate API Gateway REST API or EventBridge ApiDestination path wildcards ("-").
	PathParameterValues string `json:"pathParameterValues,omitempty" yaml:"pathParameterValues,omitempty"`

	// Key-value mapping of the query strings that need to be sent as part of request invoking the API Gateway REST API or EventBridge ApiDestination.
	QueryStringParameters map[string]string `json:"queryStringParameters,omitempty" yaml:"queryStringParameters,omitempty"`
}
