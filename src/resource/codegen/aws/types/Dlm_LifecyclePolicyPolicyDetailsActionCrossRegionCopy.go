package types

type Dlm_LifecyclePolicyPolicyDetailsActionCrossRegionCopy struct {
	// The encryption settings for the copied snapshot. See the `encryption_configuration` block. Max of 1 per action.
	EncryptionConfiguration Dlm_LifecyclePolicyPolicyDetailsActionCrossRegionCopyEncryptionConfiguration `json:"encryptionConfiguration,omitempty" yaml:"encryptionConfiguration,omitempty"`

	//
	RetainRule Dlm_LifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRule `json:"retainRule,omitempty" yaml:"retainRule,omitempty"`

	//
	Target string `json:"target,omitempty" yaml:"target,omitempty"`
}
