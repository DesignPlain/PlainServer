package types

type Appmesh_RouteSpecTcpRouteTimeout struct {
	// Idle timeout. An idle timeout bounds the amount of time that a connection may be idle.
	Idle Appmesh_RouteSpecTcpRouteTimeoutIdle `json:"idle,omitempty" yaml:"idle,omitempty"`
}
