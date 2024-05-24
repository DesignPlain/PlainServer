package types

type Appmesh_getVirtualNodeSpecListenerTimeoutHttp2 struct {
	//
	Idles []Appmesh_getVirtualNodeSpecListenerTimeoutHttp2Idle `json:"idles,omitempty" yaml:"idles,omitempty"`

	//
	PerRequests []Appmesh_getVirtualNodeSpecListenerTimeoutHttp2PerRequest `json:"perRequests,omitempty" yaml:"perRequests,omitempty"`
}
