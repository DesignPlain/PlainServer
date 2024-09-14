package types

type Appmesh_VirtualNodeSpecListenerConnectionPoolTcp struct {
	// Maximum number of outbound TCP connections Envoy can establish concurrently with all hosts in upstream cluster. Minimum value of `1`.
	MaxConnections int `json:"maxConnections,omitempty" yaml:"maxConnections,omitempty"`
}
