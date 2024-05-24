package types

type Appmesh_getRouteSpecHttpRouteTimeout struct {
	//
	Idles []Appmesh_getRouteSpecHttpRouteTimeoutIdle `json:"idles,omitempty" yaml:"idles,omitempty"`

	//
	PerRequests []Appmesh_getRouteSpecHttpRouteTimeoutPerRequest `json:"perRequests,omitempty" yaml:"perRequests,omitempty"`
}
