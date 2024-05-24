package types

type Appmesh_getVirtualGatewaySpecListenerConnectionPool struct {
	//
	Grpcs []Appmesh_getVirtualGatewaySpecListenerConnectionPoolGrpc `json:"grpcs,omitempty" yaml:"grpcs,omitempty"`

	//
	Http2s []Appmesh_getVirtualGatewaySpecListenerConnectionPoolHttp2 `json:"http2s,omitempty" yaml:"http2s,omitempty"`

	//
	Https []Appmesh_getVirtualGatewaySpecListenerConnectionPoolHttp `json:"https,omitempty" yaml:"https,omitempty"`
}
