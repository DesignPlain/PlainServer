package types

type Emrserverless_ApplicationInitialCapacityInitialCapacityConfig struct {
	// The resource configuration of the initial capacity configuration.
	WorkerConfiguration Emrserverless_ApplicationInitialCapacityInitialCapacityConfigWorkerConfiguration `json:"workerConfiguration,omitempty" yaml:"workerConfiguration,omitempty"`

	// The number of workers in the initial capacity configuration.
	WorkerCount int `json:"workerCount,omitempty" yaml:"workerCount,omitempty"`
}
