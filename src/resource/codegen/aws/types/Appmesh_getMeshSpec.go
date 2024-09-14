package types

type Appmesh_getMeshSpec struct {
	//
	ServiceDiscoveries []Appmesh_getMeshSpecServiceDiscovery `json:"serviceDiscoveries,omitempty" yaml:"serviceDiscoveries,omitempty"`

	//
	EgressFilters []Appmesh_getMeshSpecEgressFilter `json:"egressFilters,omitempty" yaml:"egressFilters,omitempty"`
}
