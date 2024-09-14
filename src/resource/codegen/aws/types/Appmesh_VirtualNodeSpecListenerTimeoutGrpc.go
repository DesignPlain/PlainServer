package types

type Appmesh_VirtualNodeSpecListenerTimeoutGrpc struct {
	// Idle timeout. An idle timeout bounds the amount of time that a connection may be idle.
	Idle Appmesh_VirtualNodeSpecListenerTimeoutGrpcIdle `json:"idle,omitempty" yaml:"idle,omitempty"`

	// Per request timeout.
	PerRequest Appmesh_VirtualNodeSpecListenerTimeoutGrpcPerRequest `json:"perRequest,omitempty" yaml:"perRequest,omitempty"`
}
