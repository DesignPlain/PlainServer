package types

type Container_ClusterNodeConfigGuestAccelerator struct {
	// Configuration for GPU sharing. Structure is documented below.
	GpuSharingConfig Container_ClusterNodeConfigGuestAcceleratorGpuSharingConfig `json:"gpuSharingConfig,omitempty" yaml:"gpuSharingConfig,omitempty"`

	// The accelerator type resource to expose to this instance. E.g. `nvidia-tesla-k80`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The number of the guest accelerator cards exposed to this instance.
	Count int `json:"count,omitempty" yaml:"count,omitempty"`

	// Configuration for auto installation of GPU driver. Structure is documented below.
	GpuDriverInstallationConfig Container_ClusterNodeConfigGuestAcceleratorGpuDriverInstallationConfig `json:"gpuDriverInstallationConfig,omitempty" yaml:"gpuDriverInstallationConfig,omitempty"`

	// Size of partitions to create on the GPU. Valid values are described in the NVIDIA mig [user guide](https://docs.nvidia.com/datacenter/tesla/mig-user-guide/#partitioning).
	GpuPartitionSize string `json:"gpuPartitionSize,omitempty" yaml:"gpuPartitionSize,omitempty"`
}
