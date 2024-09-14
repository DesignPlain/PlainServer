package types

type Container_ClusterNodePoolUpgradeSettings struct {
	// Settings for blue-green upgrade strategy. To be specified when strategy is set to BLUE_GREEN. Structure is documented below.
	BlueGreenSettings Container_ClusterNodePoolUpgradeSettingsBlueGreenSettings `json:"blueGreenSettings,omitempty" yaml:"blueGreenSettings,omitempty"`

	// The maximum number of nodes that can be created beyond the current size of the node pool during the upgrade process. To be used when strategy is set to SURGE. Default is 0.
	MaxSurge int `json:"maxSurge,omitempty" yaml:"maxSurge,omitempty"`

	// The maximum number of nodes that can be simultaneously unavailable during the upgrade process. To be used when strategy is set to SURGE. Default is 0.
	MaxUnavailable int `json:"maxUnavailable,omitempty" yaml:"maxUnavailable,omitempty"`

	// Strategy used for node pool update. Strategy can only be one of BLUE_GREEN or SURGE. The default is value is SURGE.
	Strategy string `json:"strategy,omitempty" yaml:"strategy,omitempty"`
}
