package types

type Dlm_LifecyclePolicyPolicyDetailsParameters struct {
	// Applies to AMI lifecycle policies only. Indicates whether targeted instances are rebooted when the lifecycle policy runs. `true` indicates that targeted instances are not rebooted when the policy runs. `false` indicates that target instances are rebooted when the policy runs. The default is `true` (instances are not rebooted).
	NoReboot bool `json:"noReboot,omitempty" yaml:"noReboot,omitempty"`

	// Indicates whether to exclude the root volume from snapshots created using CreateSnapshots. The default is `false`.
	ExcludeBootVolume bool `json:"excludeBootVolume,omitempty" yaml:"excludeBootVolume,omitempty"`
}
