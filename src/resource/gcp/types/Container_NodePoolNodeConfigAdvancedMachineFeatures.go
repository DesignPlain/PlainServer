package types

type Container_NodePoolNodeConfigAdvancedMachineFeatures struct {
	// The number of threads per physical core. To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed.
	ThreadsPerCore int `json:"threadsPerCore,omitempty" yaml:"threadsPerCore,omitempty"`
}
