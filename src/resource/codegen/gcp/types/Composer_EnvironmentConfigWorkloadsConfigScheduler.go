package types

type Composer_EnvironmentConfigWorkloadsConfigScheduler struct {
	// The number of schedulers.
	Count int `json:"count,omitempty" yaml:"count,omitempty"`

	// CPU request and limit for a single Airflow scheduler replica
	Cpu float64 `json:"cpu,omitempty" yaml:"cpu,omitempty"`

	// Memory (GB) request and limit for a single Airflow scheduler replica.
	MemoryGb float64 `json:"memoryGb,omitempty" yaml:"memoryGb,omitempty"`

	// Storage (GB) request and limit for a single Airflow scheduler replica.
	StorageGb float64 `json:"storageGb,omitempty" yaml:"storageGb,omitempty"`
}
