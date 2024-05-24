package types

type Appmesh_VirtualGatewaySpecListenerConnectionPool struct {
	// Connection pool information for HTTP listeners.
	Http Appmesh_VirtualGatewaySpecListenerConnectionPoolHttp `json:"http,omitempty" yaml:"http,omitempty"`

	// Connection pool information for HTTP2 listeners.
	Http2 Appmesh_VirtualGatewaySpecListenerConnectionPoolHttp2 `json:"http2,omitempty" yaml:"http2,omitempty"`

	// Connection pool information for gRPC listeners.
	Grpc Appmesh_VirtualGatewaySpecListenerConnectionPoolGrpc `json:"grpc,omitempty" yaml:"grpc,omitempty"`
}
