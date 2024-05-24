package types

type Appmesh_GatewayRouteSpecHttp2RouteActionTargetVirtualService struct {
	// Name of the virtual service that traffic is routed to. Must be between 1 and 255 characters in length.
	VirtualServiceName string `json:"virtualServiceName,omitempty" yaml:"virtualServiceName,omitempty"`
}
