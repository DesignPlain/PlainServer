package types

type Compute_InstanceAdvancedMachineFeatures struct {
	// The number of physical cores to expose to an instance. [visible cores info (VC)](https://cloud.google.com/compute/docs/instances/customize-visible-cores).
	VisibleCoreCount int `json:"visibleCoreCount,omitempty" yaml:"visibleCoreCount,omitempty"`

	// Defines whether the instance should have nested virtualization  enabled. Defaults to false.
	EnableNestedVirtualization bool `json:"enableNestedVirtualization,omitempty" yaml:"enableNestedVirtualization,omitempty"`

	// he number of threads per physical core. To disable [simultaneous multithreading (SMT)](https://cloud.google.com/compute/docs/instances/disabling-smt) set this to 1.
	ThreadsPerCore int `json:"threadsPerCore,omitempty" yaml:"threadsPerCore,omitempty"`
}
