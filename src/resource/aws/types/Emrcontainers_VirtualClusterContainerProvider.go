package types

type Emrcontainers_VirtualClusterContainerProvider struct {
	// The name of the container provider that is running your EMR Containers cluster
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Nested list containing information about the configuration of the container provider
	Info Emrcontainers_VirtualClusterContainerProviderInfo `json:"info,omitempty" yaml:"info,omitempty"`

	// The type of the container provider
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
