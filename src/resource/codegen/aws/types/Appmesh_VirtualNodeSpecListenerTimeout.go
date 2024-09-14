package types

type Appmesh_VirtualNodeSpecListenerTimeout struct {
	// Timeouts for TCP listeners.
	Tcp Appmesh_VirtualNodeSpecListenerTimeoutTcp `json:"tcp,omitempty" yaml:"tcp,omitempty"`

	// Timeouts for gRPC listeners.
	Grpc Appmesh_VirtualNodeSpecListenerTimeoutGrpc `json:"grpc,omitempty" yaml:"grpc,omitempty"`

	// Timeouts for HTTP listeners.
	Http Appmesh_VirtualNodeSpecListenerTimeoutHttp `json:"http,omitempty" yaml:"http,omitempty"`

	// Timeouts for HTTP2 listeners.
	Http2 Appmesh_VirtualNodeSpecListenerTimeoutHttp2 `json:"http2,omitempty" yaml:"http2,omitempty"`
}
