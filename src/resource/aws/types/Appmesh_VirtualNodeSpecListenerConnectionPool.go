package types

type Appmesh_VirtualNodeSpecListenerConnectionPool struct {
	// Connection pool information for gRPC listeners.
	Grpc Appmesh_VirtualNodeSpecListenerConnectionPoolGrpc `json:"grpc,omitempty" yaml:"grpc,omitempty"`

	// Connection pool information for HTTP2 listeners.
	Http2s []Appmesh_VirtualNodeSpecListenerConnectionPoolHttp2 `json:"http2s,omitempty" yaml:"http2s,omitempty"`

	// Connection pool information for HTTP listeners.
	Https []Appmesh_VirtualNodeSpecListenerConnectionPoolHttp `json:"https,omitempty" yaml:"https,omitempty"`

	// Connection pool information for TCP listeners.
	Tcps []Appmesh_VirtualNodeSpecListenerConnectionPoolTcp `json:"tcps,omitempty" yaml:"tcps,omitempty"`
}
