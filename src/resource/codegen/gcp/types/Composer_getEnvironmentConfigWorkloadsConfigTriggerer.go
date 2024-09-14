package types

type Composer_getEnvironmentConfigWorkloadsConfigTriggerer struct {
	// Memory (GB) request and limit for a single Airflow triggerer replica.
	MemoryGb float64 `json:"memoryGb,omitempty" yaml:"memoryGb,omitempty"`

	// The number of triggerers.
	Count int `json:"count,omitempty" yaml:"count,omitempty"`

	// CPU request and limit for a single Airflow triggerer replica.
	Cpu float64 `json:"cpu,omitempty" yaml:"cpu,omitempty"`
}
