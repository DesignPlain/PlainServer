package types

type Appmesh_getRouteSpec struct {
	//
	HttpRoutes []Appmesh_getRouteSpecHttpRoute `json:"httpRoutes,omitempty" yaml:"httpRoutes,omitempty"`

	//
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	//
	TcpRoutes []Appmesh_getRouteSpecTcpRoute `json:"tcpRoutes,omitempty" yaml:"tcpRoutes,omitempty"`

	//
	GrpcRoutes []Appmesh_getRouteSpecGrpcRoute `json:"grpcRoutes,omitempty" yaml:"grpcRoutes,omitempty"`

	//
	Http2Routes []Appmesh_getRouteSpecHttp2Route `json:"http2Routes,omitempty" yaml:"http2Routes,omitempty"`
}
