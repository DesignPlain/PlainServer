package types

type Appmesh_getVirtualNodeSpec struct {
	//
	ServiceDiscoveries []Appmesh_getVirtualNodeSpecServiceDiscovery `json:"serviceDiscoveries,omitempty" yaml:"serviceDiscoveries,omitempty"`

	//
	BackendDefaults []Appmesh_getVirtualNodeSpecBackendDefault `json:"backendDefaults,omitempty" yaml:"backendDefaults,omitempty"`

	//
	Backends []Appmesh_getVirtualNodeSpecBackend `json:"backends,omitempty" yaml:"backends,omitempty"`

	//
	Listeners []Appmesh_getVirtualNodeSpecListener `json:"listeners,omitempty" yaml:"listeners,omitempty"`

	//
	Loggings []Appmesh_getVirtualNodeSpecLogging `json:"loggings,omitempty" yaml:"loggings,omitempty"`
}
