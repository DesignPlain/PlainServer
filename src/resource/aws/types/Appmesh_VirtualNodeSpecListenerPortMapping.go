package types

type Appmesh_VirtualNodeSpecListenerPortMapping struct {
	// Port used for the port mapping.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Protocol used for the port mapping. Valid values are `http`, `http2`, `tcp` and `grpc`.
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
}
