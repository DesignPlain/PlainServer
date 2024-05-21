package types

type Container_getClusterNodePoolUpgradeSetting struct {
	// Update strategy for the given nodepool.
	Strategy string `json:"strategy,omitempty" yaml:"strategy,omitempty"`

	// Settings for BlueGreen node pool upgrade.
	BlueGreenSettings []Container_getClusterNodePoolUpgradeSettingBlueGreenSetting `json:"blueGreenSettings,omitempty" yaml:"blueGreenSettings,omitempty"`

	// The number of additional nodes that can be added to the node pool during an upgrade. Increasing max_surge raises the number of nodes that can be upgraded simultaneously. Can be set to 0 or greater.
	MaxSurge int `json:"maxSurge,omitempty" yaml:"maxSurge,omitempty"`

	// The number of nodes that can be simultaneously unavailable during an upgrade. Increasing max_unavailable raises the number of nodes that can be upgraded in parallel. Can be set to 0 or greater.
	MaxUnavailable int `json:"maxUnavailable,omitempty" yaml:"maxUnavailable,omitempty"`
}
