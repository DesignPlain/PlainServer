package types

type Dlm_LifecyclePolicyPolicyDetailsActionCrossRegionCopy struct {
	// The encryption settings for the copied snapshot. See the `encryption_configuration` block. Max of 1 per action.
	EncryptionConfiguration Dlm_LifecyclePolicyPolicyDetailsActionCrossRegionCopyEncryptionConfiguration `json:"encryptionConfiguration,omitempty" yaml:"encryptionConfiguration,omitempty"`

	// Specifies the retention rule for cross-Region snapshot copies. See the `retain_rule` block. Max of 1 per action.
	RetainRule Dlm_LifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRule `json:"retainRule,omitempty" yaml:"retainRule,omitempty"`

	// The target Region or the Amazon Resource Name (ARN) of the target Outpost for the snapshot copies.
	Target string `json:"target,omitempty" yaml:"target,omitempty"`
}
