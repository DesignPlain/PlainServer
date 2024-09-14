package types

type Dlm_LifecyclePolicyPolicyDetailsScheduleShareRule struct {
	// The IDs of the AWS accounts with which to share the snapshots.
	TargetAccounts []string `json:"targetAccounts,omitempty" yaml:"targetAccounts,omitempty"`

	// The period after which snapshots that are shared with other AWS accounts are automatically unshared.
	UnshareInterval int `json:"unshareInterval,omitempty" yaml:"unshareInterval,omitempty"`

	// The unit of time for the automatic unsharing interval. Valid values are `DAYS`, `WEEKS`, `MONTHS`, `YEARS`.
	UnshareIntervalUnit string `json:"unshareIntervalUnit,omitempty" yaml:"unshareIntervalUnit,omitempty"`
}
