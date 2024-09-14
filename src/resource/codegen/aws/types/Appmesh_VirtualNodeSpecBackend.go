package types

type Appmesh_VirtualNodeSpecBackend struct {
	// Virtual service to use as a backend for a virtual node.
	VirtualService Appmesh_VirtualNodeSpecBackendVirtualService `json:"virtualService,omitempty" yaml:"virtualService,omitempty"`
}
