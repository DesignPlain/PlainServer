package types

type Appmesh_VirtualServiceSpec struct {
	// App Mesh object that is acting as the provider for a virtual service. You can specify a single virtual node or virtual router.
	Provider Appmesh_VirtualServiceSpecProvider `json:"provider,omitempty" yaml:"provider,omitempty"`
}
