package types

type Emrcontainers_getVirtualClusterContainerProvider struct {
	// The name of the container provider that is running your EMR Containers cluster
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Nested list containing information about the configuration of the container provider
	Infos []Emrcontainers_getVirtualClusterContainerProviderInfo `json:"infos,omitempty" yaml:"infos,omitempty"`

	// The type of the container provider
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
