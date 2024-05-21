package types

type Composer_getEnvironmentConfigWorkloadsConfigWorker struct {
	// Storage (GB) request and limit for a single Airflow worker replica.
	StorageGb float64 `json:"storageGb,omitempty" yaml:"storageGb,omitempty"`

	// CPU request and limit for a single Airflow worker replica.
	Cpu float64 `json:"cpu,omitempty" yaml:"cpu,omitempty"`

	// Maximum number of workers for autoscaling.
	MaxCount int `json:"maxCount,omitempty" yaml:"maxCount,omitempty"`

	// Memory (GB) request and limit for a single Airflow worker replica.
	MemoryGb float64 `json:"memoryGb,omitempty" yaml:"memoryGb,omitempty"`

	// Minimum number of workers for autoscaling.
	MinCount int `json:"minCount,omitempty" yaml:"minCount,omitempty"`
}
