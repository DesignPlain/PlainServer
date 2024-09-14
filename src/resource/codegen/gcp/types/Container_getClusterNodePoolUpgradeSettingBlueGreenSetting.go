package types

type Container_getClusterNodePoolUpgradeSettingBlueGreenSetting struct {
	// Time needed after draining entire blue pool. After this period, blue pool will be cleaned up.
	NodePoolSoakDuration string `json:"nodePoolSoakDuration,omitempty" yaml:"nodePoolSoakDuration,omitempty"`

	// Standard rollout policy is the default policy for blue-green.
	StandardRolloutPolicies []Container_getClusterNodePoolUpgradeSettingBlueGreenSettingStandardRolloutPolicy `json:"standardRolloutPolicies,omitempty" yaml:"standardRolloutPolicies,omitempty"`
}
