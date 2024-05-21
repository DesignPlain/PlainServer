package types

type Compute_BackendServiceCircuitBreakers struct {
	/*
	   Maximum requests for a single backend connection. This parameter
	   is respected by both the HTTP/1.1 and HTTP/2 implementations. If
	   not specified, there is no limit. Setting this parameter to 1
	   will effectively disable keep alive.
	*/
	MaxRequestsPerConnection int `json:"maxRequestsPerConnection,omitempty" yaml:"maxRequestsPerConnection,omitempty"`

	/*
	   The maximum number of parallel retries to the backend cluster.
	   Defaults to 3.
	*/
	MaxRetries int `json:"maxRetries,omitempty" yaml:"maxRetries,omitempty"`

	/*
	   The timeout for new network connections to hosts.
	   Structure is documented below.
	*/
	ConnectTimeout Compute_BackendServiceCircuitBreakersConnectTimeout `json:"connectTimeout,omitempty" yaml:"connectTimeout,omitempty"`

	/*
	   The maximum number of connections to the backend cluster.
	   Defaults to 1024.
	*/
	MaxConnections int `json:"maxConnections,omitempty" yaml:"maxConnections,omitempty"`

	/*
	   The maximum number of pending requests to the backend cluster.
	   Defaults to 1024.
	*/
	MaxPendingRequests int `json:"maxPendingRequests,omitempty" yaml:"maxPendingRequests,omitempty"`

	/*
	   The maximum number of parallel requests to the backend cluster.
	   Defaults to 1024.
	*/
	MaxRequests int `json:"maxRequests,omitempty" yaml:"maxRequests,omitempty"`
}
