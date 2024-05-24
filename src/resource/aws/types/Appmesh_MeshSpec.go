package types

type Appmesh_MeshSpec struct {
	// Egress filter rules for the service mesh.
	EgressFilter Appmesh_MeshSpecEgressFilter `json:"egressFilter,omitempty" yaml:"egressFilter,omitempty"`
}
