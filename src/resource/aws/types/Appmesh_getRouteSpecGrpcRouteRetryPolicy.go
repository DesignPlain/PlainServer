package types

type Appmesh_getRouteSpecGrpcRouteRetryPolicy struct {
	//
	TcpRetryEvents []string `json:"tcpRetryEvents,omitempty" yaml:"tcpRetryEvents,omitempty"`

	//
	GrpcRetryEvents []string `json:"grpcRetryEvents,omitempty" yaml:"grpcRetryEvents,omitempty"`

	//
	HttpRetryEvents []string `json:"httpRetryEvents,omitempty" yaml:"httpRetryEvents,omitempty"`

	//
	MaxRetries int `json:"maxRetries,omitempty" yaml:"maxRetries,omitempty"`

	//
	PerRetryTimeouts []Appmesh_getRouteSpecGrpcRouteRetryPolicyPerRetryTimeout `json:"perRetryTimeouts,omitempty" yaml:"perRetryTimeouts,omitempty"`
}
