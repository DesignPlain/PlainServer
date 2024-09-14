package types

type Appmesh_getGatewayRouteSpec struct {
	//
	GrpcRoutes []Appmesh_getGatewayRouteSpecGrpcRoute `json:"grpcRoutes,omitempty" yaml:"grpcRoutes,omitempty"`

	//
	Http2Routes []Appmesh_getGatewayRouteSpecHttp2Route `json:"http2Routes,omitempty" yaml:"http2Routes,omitempty"`

	//
	HttpRoutes []Appmesh_getGatewayRouteSpecHttpRoute `json:"httpRoutes,omitempty" yaml:"httpRoutes,omitempty"`

	//
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`
}
