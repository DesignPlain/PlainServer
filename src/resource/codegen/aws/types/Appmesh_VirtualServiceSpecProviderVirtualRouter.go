package types

type Appmesh_VirtualServiceSpecProviderVirtualRouter struct {
	// Name of the virtual router that is acting as a service provider. Must be between 1 and 255 characters in length.
	VirtualRouterName string `json:"virtualRouterName,omitempty" yaml:"virtualRouterName,omitempty"`
}
