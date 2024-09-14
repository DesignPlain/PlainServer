package types

type Appmesh_getVirtualNodeSpecListenerTimeoutHttp2 struct {
	//
	PerRequests []Appmesh_getVirtualNodeSpecListenerTimeoutHttp2PerRequest `json:"perRequests,omitempty" yaml:"perRequests,omitempty"`

	//
	Idles []Appmesh_getVirtualNodeSpecListenerTimeoutHttp2Idle `json:"idles,omitempty" yaml:"idles,omitempty"`
}
