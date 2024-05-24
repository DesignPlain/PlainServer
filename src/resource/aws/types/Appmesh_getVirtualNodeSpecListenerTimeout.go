package types

type Appmesh_getVirtualNodeSpecListenerTimeout struct {
	//
	Http2s []Appmesh_getVirtualNodeSpecListenerTimeoutHttp2 `json:"http2s,omitempty" yaml:"http2s,omitempty"`

	//
	Https []Appmesh_getVirtualNodeSpecListenerTimeoutHttp `json:"https,omitempty" yaml:"https,omitempty"`

	//
	Tcps []Appmesh_getVirtualNodeSpecListenerTimeoutTcp `json:"tcps,omitempty" yaml:"tcps,omitempty"`

	//
	Grpcs []Appmesh_getVirtualNodeSpecListenerTimeoutGrpc `json:"grpcs,omitempty" yaml:"grpcs,omitempty"`
}
