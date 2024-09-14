package types

type Appmesh_VirtualNodeSpecListenerTimeoutHttp2 struct {
	// Per request timeout.
	PerRequest Appmesh_VirtualNodeSpecListenerTimeoutHttp2PerRequest `json:"perRequest,omitempty" yaml:"perRequest,omitempty"`

	// Idle timeout. An idle timeout bounds the amount of time that a connection may be idle.
	Idle Appmesh_VirtualNodeSpecListenerTimeoutHttp2Idle `json:"idle,omitempty" yaml:"idle,omitempty"`
}
