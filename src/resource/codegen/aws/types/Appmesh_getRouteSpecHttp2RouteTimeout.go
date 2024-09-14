package types

type Appmesh_getRouteSpecHttp2RouteTimeout struct {
	//
	Idles []Appmesh_getRouteSpecHttp2RouteTimeoutIdle `json:"idles,omitempty" yaml:"idles,omitempty"`

	//
	PerRequests []Appmesh_getRouteSpecHttp2RouteTimeoutPerRequest `json:"perRequests,omitempty" yaml:"perRequests,omitempty"`
}
