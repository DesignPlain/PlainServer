package types

type Emrserverless_ApplicationInitialCapacityInitialCapacityConfigWorkerConfiguration struct {
	// The disk requirements for every worker instance of the worker type.
	Disk string `json:"disk,omitempty" yaml:"disk,omitempty"`

	// The memory requirements for every worker instance of the worker type.
	Memory string `json:"memory,omitempty" yaml:"memory,omitempty"`

	// The CPU requirements for every worker instance of the worker type.
	Cpu string `json:"cpu,omitempty" yaml:"cpu,omitempty"`
}
