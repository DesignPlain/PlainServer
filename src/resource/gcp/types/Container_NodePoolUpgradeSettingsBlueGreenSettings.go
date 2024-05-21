package types

type Container_NodePoolUpgradeSettingsBlueGreenSettings struct {
	/*
	   Time needed after draining the entire blue pool.
	   After this period, the blue pool will be cleaned up.
	*/
	NodePoolSoakDuration string `json:"nodePoolSoakDuration,omitempty" yaml:"nodePoolSoakDuration,omitempty"`

	// Specifies the standard policy settings for blue-green upgrades.
	StandardRolloutPolicy Container_NodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy `json:"standardRolloutPolicy,omitempty" yaml:"standardRolloutPolicy,omitempty"`
}
