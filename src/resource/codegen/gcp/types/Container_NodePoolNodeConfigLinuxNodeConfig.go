package types

type Container_NodePoolNodeConfigLinuxNodeConfig struct {
	// cgroupMode specifies the cgroup mode to be used on the node.
	CgroupMode string `json:"cgroupMode,omitempty" yaml:"cgroupMode,omitempty"`

	// The Linux kernel parameters to be applied to the nodes and all pods running on the nodes.
	Sysctls map[string]string `json:"sysctls,omitempty" yaml:"sysctls,omitempty"`
}
