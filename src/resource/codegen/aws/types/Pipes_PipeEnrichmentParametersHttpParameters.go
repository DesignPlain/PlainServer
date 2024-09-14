package types

type Pipes_PipeEnrichmentParametersHttpParameters struct {
	//
	HeaderParameters map[string]string `json:"headerParameters,omitempty" yaml:"headerParameters,omitempty"`

	//
	PathParameterValues string `json:"pathParameterValues,omitempty" yaml:"pathParameterValues,omitempty"`

	//
	QueryStringParameters map[string]string `json:"queryStringParameters,omitempty" yaml:"queryStringParameters,omitempty"`
}
