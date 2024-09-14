package types

type Composer_EnvironmentConfigWorkloadsConfigDagProcessor struct {
	// CPU request and limit for DAG processor.
	Cpu float64 `json:"cpu,omitempty" yaml:"cpu,omitempty"`

	// Memory (GB) request and limit for DAG processor.
	MemoryGb float64 `json:"memoryGb,omitempty" yaml:"memoryGb,omitempty"`

	// Storage (GB) request and limit for DAG processor.
	StorageGb float64 `json:"storageGb,omitempty" yaml:"storageGb,omitempty"`
}
