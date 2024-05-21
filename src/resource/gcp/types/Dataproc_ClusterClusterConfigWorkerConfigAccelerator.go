package types

type Dataproc_ClusterClusterConfigWorkerConfigAccelerator struct {
	/*
	   The number of the accelerator cards of this type exposed to this instance. Often restricted to one of `1`, `2`, `4`, or `8`.


	   - - -
	*/
	AcceleratorCount int `json:"acceleratorCount,omitempty" yaml:"acceleratorCount,omitempty"`

	// The short name of the accelerator type to expose to this instance. For example, `nvidia-tesla-k80`.
	AcceleratorType string `json:"acceleratorType,omitempty" yaml:"acceleratorType,omitempty"`
}
