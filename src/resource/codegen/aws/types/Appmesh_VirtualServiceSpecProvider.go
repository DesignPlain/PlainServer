package types

type Appmesh_VirtualServiceSpecProvider struct {
	// Virtual node associated with a virtual service.
	VirtualNode Appmesh_VirtualServiceSpecProviderVirtualNode `json:"virtualNode,omitempty" yaml:"virtualNode,omitempty"`

	// Virtual router associated with a virtual service.
	VirtualRouter Appmesh_VirtualServiceSpecProviderVirtualRouter `json:"virtualRouter,omitempty" yaml:"virtualRouter,omitempty"`
}
