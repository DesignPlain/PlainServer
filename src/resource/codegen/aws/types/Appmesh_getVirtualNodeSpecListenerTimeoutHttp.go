package types

type Appmesh_getVirtualNodeSpecListenerTimeoutHttp struct {
	//
	Idles []Appmesh_getVirtualNodeSpecListenerTimeoutHttpIdle `json:"idles,omitempty" yaml:"idles,omitempty"`

	//
	PerRequests []Appmesh_getVirtualNodeSpecListenerTimeoutHttpPerRequest `json:"perRequests,omitempty" yaml:"perRequests,omitempty"`
}
