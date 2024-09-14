package types

type Appmesh_getVirtualNodeSpecListenerTimeoutGrpc struct {
	//
	PerRequests []Appmesh_getVirtualNodeSpecListenerTimeoutGrpcPerRequest `json:"perRequests,omitempty" yaml:"perRequests,omitempty"`

	//
	Idles []Appmesh_getVirtualNodeSpecListenerTimeoutGrpcIdle `json:"idles,omitempty" yaml:"idles,omitempty"`
}
