package types

type Composer_getEnvironmentConfigWorkloadsConfigWebServer struct {
	// CPU request and limit for Airflow web server.
	Cpu float64 `json:"cpu,omitempty" yaml:"cpu,omitempty"`

	// Memory (GB) request and limit for Airflow web server.
	MemoryGb float64 `json:"memoryGb,omitempty" yaml:"memoryGb,omitempty"`

	// Storage (GB) request and limit for Airflow web server.
	StorageGb float64 `json:"storageGb,omitempty" yaml:"storageGb,omitempty"`
}
