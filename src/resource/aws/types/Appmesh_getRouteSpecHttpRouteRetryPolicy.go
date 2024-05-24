package types

type Appmesh_getRouteSpecHttpRouteRetryPolicy struct {
	//
	HttpRetryEvents []string `json:"httpRetryEvents,omitempty" yaml:"httpRetryEvents,omitempty"`

	//
	MaxRetries int `json:"maxRetries,omitempty" yaml:"maxRetries,omitempty"`

	//
	PerRetryTimeouts []Appmesh_getRouteSpecHttpRouteRetryPolicyPerRetryTimeout `json:"perRetryTimeouts,omitempty" yaml:"perRetryTimeouts,omitempty"`

	//
	TcpRetryEvents []string `json:"tcpRetryEvents,omitempty" yaml:"tcpRetryEvents,omitempty"`
}
