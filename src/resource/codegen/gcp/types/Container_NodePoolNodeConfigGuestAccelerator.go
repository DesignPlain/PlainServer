package types

type Container_NodePoolNodeConfigGuestAccelerator struct {
	// The number of the accelerator cards exposed to an instance.
	Count int `json:"count,omitempty" yaml:"count,omitempty"`

	// Configuration for auto installation of GPU driver.
	GpuDriverInstallationConfig Container_NodePoolNodeConfigGuestAcceleratorGpuDriverInstallationConfig `json:"gpuDriverInstallationConfig,omitempty" yaml:"gpuDriverInstallationConfig,omitempty"`

	// Size of partitions to create on the GPU. Valid values are described in the NVIDIA mig user guide (https://docs.nvidia.com/datacenter/tesla/mig-user-guide/#partitioning)
	GpuPartitionSize string `json:"gpuPartitionSize,omitempty" yaml:"gpuPartitionSize,omitempty"`

	// Configuration for GPU sharing.
	GpuSharingConfig Container_NodePoolNodeConfigGuestAcceleratorGpuSharingConfig `json:"gpuSharingConfig,omitempty" yaml:"gpuSharingConfig,omitempty"`

	/*
	   The type of the policy. Supports a single value: COMPACT.
	   Specifying COMPACT placement policy type places node pool's nodes in a closer
	   physical proximity in order to reduce network latency between nodes.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
