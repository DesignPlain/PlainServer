package types

type Emrserverless_ApplicationInteractiveConfiguration struct {
	// Enables an Apache Livy endpoint that you can connect to and run interactive jobs.
	LivyEndpointEnabled bool `json:"livyEndpointEnabled,omitempty" yaml:"livyEndpointEnabled,omitempty"`

	// Enables you to connect an application to Amazon EMR Studio to run interactive workloads in a notebook.
	StudioEnabled bool `json:"studioEnabled,omitempty" yaml:"studioEnabled,omitempty"`
}
