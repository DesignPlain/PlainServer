package types

type Appmesh_VirtualNodeSpecListenerTimeoutTcp struct {
	// Idle timeout. An idle timeout bounds the amount of time that a connection may be idle.
	Idle Appmesh_VirtualNodeSpecListenerTimeoutTcpIdle `json:"idle,omitempty" yaml:"idle,omitempty"`
}
