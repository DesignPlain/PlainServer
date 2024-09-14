package types

type Appmesh_RouteSpec struct {
	// GRPC routing information for the route.
	GrpcRoute Appmesh_RouteSpecGrpcRoute `json:"grpcRoute,omitempty" yaml:"grpcRoute,omitempty"`

	// HTTP/2 routing information for the route.
	Http2Route Appmesh_RouteSpecHttp2Route `json:"http2Route,omitempty" yaml:"http2Route,omitempty"`

	// HTTP routing information for the route.
	HttpRoute Appmesh_RouteSpecHttpRoute `json:"httpRoute,omitempty" yaml:"httpRoute,omitempty"`

	/*
	   Priority for the route, between `0` and `1000`.
	   Routes are matched based on the specified value, where `0` is the highest priority.
	*/
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// TCP routing information for the route.
	TcpRoute Appmesh_RouteSpecTcpRoute `json:"tcpRoute,omitempty" yaml:"tcpRoute,omitempty"`
}
