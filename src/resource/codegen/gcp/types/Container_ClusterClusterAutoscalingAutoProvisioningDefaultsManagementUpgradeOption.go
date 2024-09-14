package types

type Container_ClusterClusterAutoscalingAutoProvisioningDefaultsManagementUpgradeOption struct {
	// This field is set when upgrades are about to commence with the approximate start time for the upgrades, in RFC3339 text format.
	AutoUpgradeStartTime string `json:"autoUpgradeStartTime,omitempty" yaml:"autoUpgradeStartTime,omitempty"`

	// Description of the cluster.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
