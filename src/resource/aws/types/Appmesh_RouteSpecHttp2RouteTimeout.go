package types

type Appmesh_RouteSpecHttp2RouteTimeout struct {
	// Idle timeout. An idle timeout bounds the amount of time that a connection may be idle.
	Idle Appmesh_RouteSpecHttp2RouteTimeoutIdle `json:"idle,omitempty" yaml:"idle,omitempty"`

	// Per request timeout.
	PerRequest Appmesh_RouteSpecHttp2RouteTimeoutPerRequest `json:"perRequest,omitempty" yaml:"perRequest,omitempty"`
}
