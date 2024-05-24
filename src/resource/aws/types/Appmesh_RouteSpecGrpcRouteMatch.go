package types

type Appmesh_RouteSpecGrpcRouteMatch struct {
	// Method name to match from the request. If you specify a name, you must also specify a `service_name`.
	MethodName string `json:"methodName,omitempty" yaml:"methodName,omitempty"`

	// The port number to match from the request.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Header value sent by the client must begin with the specified characters.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// Fully qualified domain name for the service to match from the request.
	ServiceName string `json:"serviceName,omitempty" yaml:"serviceName,omitempty"`

	// Data to match from the gRPC request.
	Metadatas []Appmesh_RouteSpecGrpcRouteMatchMetadata `json:"metadatas,omitempty" yaml:"metadatas,omitempty"`
}
