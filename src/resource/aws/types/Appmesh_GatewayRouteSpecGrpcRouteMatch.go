package types

type Appmesh_GatewayRouteSpecGrpcRouteMatch struct {
	// The port number to match from the request.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Fully qualified domain name for the service to match from the request.
	ServiceName string `json:"serviceName,omitempty" yaml:"serviceName,omitempty"`
}
