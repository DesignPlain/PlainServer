package types

type Appmesh_VirtualNodeSpecBackendVirtualService struct {
	// Client policy for the backend.
	ClientPolicy Appmesh_VirtualNodeSpecBackendVirtualServiceClientPolicy `json:"clientPolicy,omitempty" yaml:"clientPolicy,omitempty"`

	// Name of the virtual service that is acting as a virtual node backend. Must be between 1 and 255 characters in length.
	VirtualServiceName string `json:"virtualServiceName,omitempty" yaml:"virtualServiceName,omitempty"`
}
