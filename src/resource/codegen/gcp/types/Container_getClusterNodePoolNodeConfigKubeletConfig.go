package types

type Container_getClusterNodePoolNodeConfigKubeletConfig struct {
	// Enable CPU CFS quota enforcement for containers that specify CPU limits.
	CpuCfsQuota bool `json:"cpuCfsQuota,omitempty" yaml:"cpuCfsQuota,omitempty"`

	// Set the CPU CFS quota period value 'cpu.cfs_period_us'.
	CpuCfsQuotaPeriod string `json:"cpuCfsQuotaPeriod,omitempty" yaml:"cpuCfsQuotaPeriod,omitempty"`

	// Control the CPU management policy on the node.
	CpuManagerPolicy string `json:"cpuManagerPolicy,omitempty" yaml:"cpuManagerPolicy,omitempty"`

	// Controls the maximum number of processes allowed to run in a pod.
	PodPidsLimit int `json:"podPidsLimit,omitempty" yaml:"podPidsLimit,omitempty"`
}
