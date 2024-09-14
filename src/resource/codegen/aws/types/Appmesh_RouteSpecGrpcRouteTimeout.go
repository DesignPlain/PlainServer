package types

type Appmesh_RouteSpecGrpcRouteTimeout struct {
	// Idle timeout. An idle timeout bounds the amount of time that a connection may be idle.
	Idle Appmesh_RouteSpecGrpcRouteTimeoutIdle `json:"idle,omitempty" yaml:"idle,omitempty"`

	// Per request timeout.
	PerRequest Appmesh_RouteSpecGrpcRouteTimeoutPerRequest `json:"perRequest,omitempty" yaml:"perRequest,omitempty"`
}
