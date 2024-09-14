package types

type Appmesh_getRouteSpecGrpcRouteTimeout struct {
	//
	Idles []Appmesh_getRouteSpecGrpcRouteTimeoutIdle `json:"idles,omitempty" yaml:"idles,omitempty"`

	//
	PerRequests []Appmesh_getRouteSpecGrpcRouteTimeoutPerRequest `json:"perRequests,omitempty" yaml:"perRequests,omitempty"`
}
