package types

type Container_getClusterClusterAutoscalingAutoProvisioningDefaultManagementUpgradeOption struct {
	// This field is set when upgrades are about to commence with the description of the upgrade.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// This field is set when upgrades are about to commence with the approximate start time for the upgrades, in RFC3339 text format.
	AutoUpgradeStartTime string `json:"autoUpgradeStartTime,omitempty" yaml:"autoUpgradeStartTime,omitempty"`
}
