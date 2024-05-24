package types

type Appmesh_getRouteSpecHttp2RouteRetryPolicy struct {
	//
	PerRetryTimeouts []Appmesh_getRouteSpecHttp2RouteRetryPolicyPerRetryTimeout `json:"perRetryTimeouts,omitempty" yaml:"perRetryTimeouts,omitempty"`

	//
	TcpRetryEvents []string `json:"tcpRetryEvents,omitempty" yaml:"tcpRetryEvents,omitempty"`

	//
	HttpRetryEvents []string `json:"httpRetryEvents,omitempty" yaml:"httpRetryEvents,omitempty"`

	//
	MaxRetries int `json:"maxRetries,omitempty" yaml:"maxRetries,omitempty"`
}
