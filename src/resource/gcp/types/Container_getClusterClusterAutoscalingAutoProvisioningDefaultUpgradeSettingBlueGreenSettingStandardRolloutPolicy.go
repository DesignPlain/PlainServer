package types

type Container_getClusterClusterAutoscalingAutoProvisioningDefaultUpgradeSettingBlueGreenSettingStandardRolloutPolicy struct {
	// Number of blue nodes to drain in a batch.
	BatchNodeCount int `json:"batchNodeCount,omitempty" yaml:"batchNodeCount,omitempty"`

	// Percentage of the bool pool nodes to drain in a batch. The range of this field should be (0.0, 1.0].
	BatchPercentage float64 `json:"batchPercentage,omitempty" yaml:"batchPercentage,omitempty"`

	/*
	   Soak time after each batch gets drained.

	   																			A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	*/
	BatchSoakDuration string `json:"batchSoakDuration,omitempty" yaml:"batchSoakDuration,omitempty"`
}
