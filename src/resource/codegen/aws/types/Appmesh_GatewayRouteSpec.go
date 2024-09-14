package types

type Appmesh_GatewayRouteSpec struct {
	// Specification of an HTTP/2 gateway route.
	Http2Route Appmesh_GatewayRouteSpecHttp2Route `json:"http2Route,omitempty" yaml:"http2Route,omitempty"`

	// Specification of an HTTP gateway route.
	HttpRoute Appmesh_GatewayRouteSpecHttpRoute `json:"httpRoute,omitempty" yaml:"httpRoute,omitempty"`

	// Priority for the gateway route, between `0` and `1000`.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// Specification of a gRPC gateway route.
	GrpcRoute Appmesh_GatewayRouteSpecGrpcRoute `json:"grpcRoute,omitempty" yaml:"grpcRoute,omitempty"`
}
