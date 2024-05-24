package types

type Appmesh_VirtualServiceSpecProviderVirtualNode struct {
	// Name of the virtual node that is acting as a service provider. Must be between 1 and 255 characters in length.
	VirtualNodeName string `json:"virtualNodeName,omitempty" yaml:"virtualNodeName,omitempty"`
}
