package types

type Container_getClusterClusterAutoscalingAutoProvisioningDefaultManagement struct {
	// Specifies whether the node auto-repair is enabled for the node pool. If enabled, the nodes in this node pool will be monitored and, if they fail health checks too many times, an automatic repair action will be triggered.
	AutoRepair bool `json:"autoRepair,omitempty" yaml:"autoRepair,omitempty"`

	// Specifies whether node auto-upgrade is enabled for the node pool. If enabled, node auto-upgrade helps keep the nodes in your node pool up to date with the latest release version of Kubernetes.
	AutoUpgrade bool `json:"autoUpgrade,omitempty" yaml:"autoUpgrade,omitempty"`

	// Specifies the Auto Upgrade knobs for the node pool.
	UpgradeOptions []Container_getClusterClusterAutoscalingAutoProvisioningDefaultManagementUpgradeOption `json:"upgradeOptions,omitempty" yaml:"upgradeOptions,omitempty"`
}
