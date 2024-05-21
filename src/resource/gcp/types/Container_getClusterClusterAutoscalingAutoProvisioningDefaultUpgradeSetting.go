package types

type Container_getClusterClusterAutoscalingAutoProvisioningDefaultUpgradeSetting struct {
	// Settings for blue-green upgrade strategy.
	BlueGreenSettings []Container_getClusterClusterAutoscalingAutoProvisioningDefaultUpgradeSettingBlueGreenSetting `json:"blueGreenSettings,omitempty" yaml:"blueGreenSettings,omitempty"`

	// The maximum number of nodes that can be created beyond the current size of the node pool during the upgrade process.
	MaxSurge int `json:"maxSurge,omitempty" yaml:"maxSurge,omitempty"`

	// The maximum number of nodes that can be simultaneously unavailable during the upgrade process.
	MaxUnavailable int `json:"maxUnavailable,omitempty" yaml:"maxUnavailable,omitempty"`

	// Update strategy of the node pool.
	Strategy string `json:"strategy,omitempty" yaml:"strategy,omitempty"`
}
