package types

type Cloudwatch_EventTargetHttpTarget struct {
	// Enables you to specify HTTP headers to add to the request.
	HeaderParameters map[string]string `json:"headerParameters,omitempty" yaml:"headerParameters,omitempty"`

	// The list of values that correspond sequentially to any path variables in your endpoint ARN (for example `arn:aws:execute-api:us-east-1:123456:myapi/-/POST/pets/-`).
	PathParameterValues []string `json:"pathParameterValues,omitempty" yaml:"pathParameterValues,omitempty"`

	// Represents keys/values of query string parameters that are appended to the invoked endpoint.
	QueryStringParameters map[string]string `json:"queryStringParameters,omitempty" yaml:"queryStringParameters,omitempty"`
}
