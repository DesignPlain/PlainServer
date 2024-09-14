package types

type Pipes_PipeEnrichmentParameters struct {
	// Contains the HTTP parameters to use when the target is a API Gateway REST endpoint or EventBridge ApiDestination. If you specify an API Gateway REST API or EventBridge ApiDestination as a target, you can use this parameter to specify headers, path parameters, and query string keys/values as part of your target invoking request. If you're using ApiDestinations, the corresponding Connection can also have these values configured. In case of any conflicting keys, values from the Connection take precedence. Detailed below.
	HttpParameters Pipes_PipeEnrichmentParametersHttpParameters `json:"httpParameters,omitempty" yaml:"httpParameters,omitempty"`

	// Valid JSON text passed to the target. In this case, nothing from the event itself is passed to the target. Maximum length of 8192 characters.
	InputTemplate string `json:"inputTemplate,omitempty" yaml:"inputTemplate,omitempty"`
}
