package types

type Appmesh_VirtualRouterSpecListener struct {
	// Port mapping information for the listener.
	PortMapping Appmesh_VirtualRouterSpecListenerPortMapping `json:"portMapping,omitempty" yaml:"portMapping,omitempty"`
}
