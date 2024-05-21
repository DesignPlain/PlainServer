package types

type Container_ClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsBlueGreenSettings struct {
	// Time needed after draining entire blue pool. After this period, blue pool will be cleaned up. A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	NodePoolSoakDuration string `json:"nodePoolSoakDuration,omitempty" yaml:"nodePoolSoakDuration,omitempty"`

	// Standard policy for the blue-green upgrade. To be specified when strategy is set to BLUE_GREEN. Structure is documented below.
	StandardRolloutPolicy Container_ClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy `json:"standardRolloutPolicy,omitempty" yaml:"standardRolloutPolicy,omitempty"`
}
