package types

type Appmesh_getVirtualNodeSpecListenerConnectionPool struct {
	//
	Http2s []Appmesh_getVirtualNodeSpecListenerConnectionPoolHttp2 `json:"http2s,omitempty" yaml:"http2s,omitempty"`

	//
	Https []Appmesh_getVirtualNodeSpecListenerConnectionPoolHttp `json:"https,omitempty" yaml:"https,omitempty"`

	//
	Tcps []Appmesh_getVirtualNodeSpecListenerConnectionPoolTcp `json:"tcps,omitempty" yaml:"tcps,omitempty"`

	//
	Grpcs []Appmesh_getVirtualNodeSpecListenerConnectionPoolGrpc `json:"grpcs,omitempty" yaml:"grpcs,omitempty"`
}
