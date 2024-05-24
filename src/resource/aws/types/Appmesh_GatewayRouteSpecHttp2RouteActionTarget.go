package types

type Appmesh_GatewayRouteSpecHttp2RouteActionTarget struct {
	// The port number that corresponds to the target for Virtual Service provider port. This is required when the provider (router or node) of the Virtual Service has multiple listeners.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Virtual service gateway route target.
	VirtualService Appmesh_GatewayRouteSpecHttp2RouteActionTargetVirtualService `json:"virtualService,omitempty" yaml:"virtualService,omitempty"`
}
