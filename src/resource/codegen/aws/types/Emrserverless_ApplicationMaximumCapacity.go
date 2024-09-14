package types

type Emrserverless_ApplicationMaximumCapacity struct {
	// The maximum allowed CPU for an application.
	Cpu string `json:"cpu,omitempty" yaml:"cpu,omitempty"`

	// The maximum allowed disk for an application.
	Disk string `json:"disk,omitempty" yaml:"disk,omitempty"`

	// The maximum allowed resources for an application.
	Memory string `json:"memory,omitempty" yaml:"memory,omitempty"`
}
