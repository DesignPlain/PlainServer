package types

type Appmesh_VirtualNodeSpecListenerTimeoutHttp struct {
	// Idle timeout. An idle timeout bounds the amount of time that a connection may be idle.
	Idle Appmesh_VirtualNodeSpecListenerTimeoutHttpIdle `json:"idle,omitempty" yaml:"idle,omitempty"`

	// Per request timeout.
	PerRequest Appmesh_VirtualNodeSpecListenerTimeoutHttpPerRequest `json:"perRequest,omitempty" yaml:"perRequest,omitempty"`
}
