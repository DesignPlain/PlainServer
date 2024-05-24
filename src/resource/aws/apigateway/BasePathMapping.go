package apigateway

type BasePathMapping struct {
	// Already-registered domain name to connect the API to.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// ID of the API to connect.
	RestApi string `json:"restApi,omitempty" yaml:"restApi,omitempty"`

	// Name of a specific deployment stage to expose at the given path. If omitted, callers may select any stage by including its name as a path element after the base path.
	StageName string `json:"stageName,omitempty" yaml:"stageName,omitempty"`

	// Path segment that must be prepended to the path when accessing the API via this mapping. If omitted, the API is exposed at the root of the given domain.
	BasePath string `json:"basePath,omitempty" yaml:"basePath,omitempty"`
}
