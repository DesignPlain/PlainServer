package types

type Appmesh_VirtualGatewaySpecListenerConnectionPoolHttp2 struct {
	// Maximum number of inflight requests Envoy can concurrently support across hosts in upstream cluster. Minimum value of `1`.
	MaxRequests int `json:"maxRequests,omitempty" yaml:"maxRequests,omitempty"`
}
