package types

type Appmesh_getVirtualNodeSpecListenerTimeoutGrpc struct {
	//
	Idles []Appmesh_getVirtualNodeSpecListenerTimeoutGrpcIdle `json:"idles,omitempty" yaml:"idles,omitempty"`

	//
	PerRequests []Appmesh_getVirtualNodeSpecListenerTimeoutGrpcPerRequest `json:"perRequests,omitempty" yaml:"perRequests,omitempty"`
}
