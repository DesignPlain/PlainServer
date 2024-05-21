package types

type Container_getClusterNodePoolUpgradeSettingBlueGreenSettingStandardRolloutPolicy struct {
	// Number of blue nodes to drain in a batch.
	BatchNodeCount int `json:"batchNodeCount,omitempty" yaml:"batchNodeCount,omitempty"`

	// Percentage of the blue pool nodes to drain in a batch.
	BatchPercentage float64 `json:"batchPercentage,omitempty" yaml:"batchPercentage,omitempty"`

	// Soak time after each batch gets drained.
	BatchSoakDuration string `json:"batchSoakDuration,omitempty" yaml:"batchSoakDuration,omitempty"`
}
