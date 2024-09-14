package types

type Container_ClusterNodeConfigGuestAcceleratorGpuSharingConfig struct {
	// The maximum number of containers that can share a GPU.
	MaxSharedClientsPerGpu int `json:"maxSharedClientsPerGpu,omitempty" yaml:"maxSharedClientsPerGpu,omitempty"`

	/*
	   The type of GPU sharing strategy to enable on the GPU node.
	   Accepted values are:
	   - `"TIME_SHARING"`: Allow multiple containers to have [time-shared](https://cloud.google.com/kubernetes-engine/docs/concepts/timesharing-gpus) access to a single GPU device.
	*/
	GpuSharingStrategy string `json:"gpuSharingStrategy,omitempty" yaml:"gpuSharingStrategy,omitempty"`
}
