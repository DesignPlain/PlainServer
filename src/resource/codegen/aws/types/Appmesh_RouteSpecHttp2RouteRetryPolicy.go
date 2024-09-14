package types

type Appmesh_RouteSpecHttp2RouteRetryPolicy struct {
	/*
	   List of HTTP retry events.
	   Valid values: `client-error` (HTTP status code 409), `gateway-error` (HTTP status codes 502, 503, and 504), `server-error` (HTTP status codes 500, 501, 502, 503, 504, 505, 506, 507, 508, 510, and 511), `stream-error` (retry on refused stream).
	*/
	HttpRetryEvents []string `json:"httpRetryEvents,omitempty" yaml:"httpRetryEvents,omitempty"`

	// Maximum number of retries.
	MaxRetries int `json:"maxRetries,omitempty" yaml:"maxRetries,omitempty"`

	// Per-retry timeout.
	PerRetryTimeout Appmesh_RouteSpecHttp2RouteRetryPolicyPerRetryTimeout `json:"perRetryTimeout,omitempty" yaml:"perRetryTimeout,omitempty"`

	/*
	   List of TCP retry events. The only valid value is `connection-error`.

	   You must specify at least one value for `http_retry_events`, or at least one value for `tcp_retry_events`.
	*/
	TcpRetryEvents []string `json:"tcpRetryEvents,omitempty" yaml:"tcpRetryEvents,omitempty"`
}
