package apigateway

type Resource struct {
	// ID of the parent API resource
	ParentId string `json:"parentId,omitempty" yaml:"parentId,omitempty"`

	// Last path segment of this API resource.
	PathPart string `json:"pathPart,omitempty" yaml:"pathPart,omitempty"`

	// ID of the associated REST API
	RestApi string `json:"restApi,omitempty" yaml:"restApi,omitempty"`
}
