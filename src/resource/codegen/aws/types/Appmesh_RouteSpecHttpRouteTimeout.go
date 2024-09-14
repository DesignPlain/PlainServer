package types

type Appmesh_RouteSpecHttpRouteTimeout struct {
	// Per request timeout.
	PerRequest Appmesh_RouteSpecHttpRouteTimeoutPerRequest `json:"perRequest,omitempty" yaml:"perRequest,omitempty"`

	// Idle timeout. An idle timeout bounds the amount of time that a connection may be idle.
	Idle Appmesh_RouteSpecHttpRouteTimeoutIdle `json:"idle,omitempty" yaml:"idle,omitempty"`
}
