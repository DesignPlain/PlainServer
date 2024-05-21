package types

type Container_getClusterNodePoolNodeConfigGuestAccelerator struct {
	// Configuration for GPU sharing.
	GpuSharingConfigs []Container_getClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfig `json:"gpuSharingConfigs,omitempty" yaml:"gpuSharingConfigs,omitempty"`

	// The accelerator type resource name.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The number of the accelerator cards exposed to an instance.
	Count int `json:"count,omitempty" yaml:"count,omitempty"`

	// Configuration for auto installation of GPU driver.
	GpuDriverInstallationConfigs []Container_getClusterNodePoolNodeConfigGuestAcceleratorGpuDriverInstallationConfig `json:"gpuDriverInstallationConfigs,omitempty" yaml:"gpuDriverInstallationConfigs,omitempty"`

	// Size of partitions to create on the GPU. Valid values are described in the NVIDIA mig user guide (https://docs.nvidia.com/datacenter/tesla/mig-user-guide/#partitioning)
	GpuPartitionSize string `json:"gpuPartitionSize,omitempty" yaml:"gpuPartitionSize,omitempty"`
}
