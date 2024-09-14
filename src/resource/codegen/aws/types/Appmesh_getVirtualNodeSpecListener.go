package types

type Appmesh_getVirtualNodeSpecListener struct {
	//
	ConnectionPools []Appmesh_getVirtualNodeSpecListenerConnectionPool `json:"connectionPools,omitempty" yaml:"connectionPools,omitempty"`

	//
	HealthChecks []Appmesh_getVirtualNodeSpecListenerHealthCheck `json:"healthChecks,omitempty" yaml:"healthChecks,omitempty"`

	//
	OutlierDetections []Appmesh_getVirtualNodeSpecListenerOutlierDetection `json:"outlierDetections,omitempty" yaml:"outlierDetections,omitempty"`

	//
	PortMappings []Appmesh_getVirtualNodeSpecListenerPortMapping `json:"portMappings,omitempty" yaml:"portMappings,omitempty"`

	//
	Timeouts []Appmesh_getVirtualNodeSpecListenerTimeout `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	//
	Tls []Appmesh_getVirtualNodeSpecListenerTl `json:"tls,omitempty" yaml:"tls,omitempty"`
}
