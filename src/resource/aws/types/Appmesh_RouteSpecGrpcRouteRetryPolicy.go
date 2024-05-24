package types

type Appmesh_RouteSpecGrpcRouteRetryPolicy struct {
	/*
	   List of gRPC retry events.
	   Valid values: `cancelled`, `deadline-exceeded`, `internal`, `resource-exhausted`, `unavailable`.
	*/
	GrpcRetryEvents []string `json:"grpcRetryEvents,omitempty" yaml:"grpcRetryEvents,omitempty"`

	/*
	   List of HTTP retry events.
	   Valid values: `client-error` (HTTP status code 409), `gateway-error` (HTTP status codes 502, 503, and 504), `server-error` (HTTP status codes 500, 501, 502, 503, 504, 505, 506, 507, 508, 510, and 511), `stream-error` (retry on refused stream).
	*/
	HttpRetryEvents []string `json:"httpRetryEvents,omitempty" yaml:"httpRetryEvents,omitempty"`

	// Maximum number of retries.
	MaxRetries int `json:"maxRetries,omitempty" yaml:"maxRetries,omitempty"`

	// Per-retry timeout.
	PerRetryTimeout Appmesh_RouteSpecGrpcRouteRetryPolicyPerRetryTimeout `json:"perRetryTimeout,omitempty" yaml:"perRetryTimeout,omitempty"`

	// List of TCP retry events. The only valid value is `connection-error`.
	TcpRetryEvents []string `json:"tcpRetryEvents,omitempty" yaml:"tcpRetryEvents,omitempty"`
}
