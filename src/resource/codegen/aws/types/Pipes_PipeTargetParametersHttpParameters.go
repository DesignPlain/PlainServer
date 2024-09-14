package types

type Pipes_PipeTargetParametersHttpParameters struct {
	//
	PathParameterValues string `json:"pathParameterValues,omitempty" yaml:"pathParameterValues,omitempty"`

	//
	QueryStringParameters map[string]string `json:"queryStringParameters,omitempty" yaml:"queryStringParameters,omitempty"`

	//
	HeaderParameters map[string]string `json:"headerParameters,omitempty" yaml:"headerParameters,omitempty"`
}
