package types

type Dlm_LifecyclePolicyPolicyDetailsScheduleShareRule struct {
	// The IDs of the AWS accounts with which to share the snapshots.
	TargetAccounts []string `json:"targetAccounts,omitempty" yaml:"targetAccounts,omitempty"`

	//
	UnshareInterval int `json:"unshareInterval,omitempty" yaml:"unshareInterval,omitempty"`

	//
	UnshareIntervalUnit string `json:"unshareIntervalUnit,omitempty" yaml:"unshareIntervalUnit,omitempty"`
}
