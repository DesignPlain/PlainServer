package types

type Appmesh_VirtualNodeSpecListenerConnectionPoolHttp struct {
	// Maximum number of outbound TCP connections Envoy can establish concurrently with all hosts in upstream cluster. Minimum value of `1`.
	MaxConnections int `json:"maxConnections,omitempty" yaml:"maxConnections,omitempty"`

	/*
	   Number of overflowing requests after `max_connections` Envoy will queue to upstream cluster. Minimum value of `1`.

	   The `http2` connection pool object supports the following:
	*/
	MaxPendingRequests int `json:"maxPendingRequests,omitempty" yaml:"maxPendingRequests,omitempty"`
}
