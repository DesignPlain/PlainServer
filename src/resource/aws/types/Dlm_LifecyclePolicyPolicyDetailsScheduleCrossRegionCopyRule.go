package types

type Dlm_LifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRule struct {
	// The Amazon Resource Name (ARN) of the AWS KMS customer master key (CMK) to use for EBS encryption. If this argument is not specified, the default KMS key for the account is used.
	CmkArn string `json:"cmkArn,omitempty" yaml:"cmkArn,omitempty"`

	// Whether to copy all user-defined tags from the source snapshot to the cross-region snapshot copy.
	CopyTags bool `json:"copyTags,omitempty" yaml:"copyTags,omitempty"`

	// The AMI deprecation rule for cross-Region AMI copies created by the rule. See the `deprecate_rule` block.
	DeprecateRule Dlm_LifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRule `json:"deprecateRule,omitempty" yaml:"deprecateRule,omitempty"`

	// To encrypt a copy of an unencrypted snapshot if encryption by default is not enabled, enable encryption using this parameter. Copies of encrypted snapshots are encrypted, even if this parameter is false or if encryption by default is not enabled.
	Encrypted bool `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`

	// Specifies the retention rule for cross-Region snapshot copies. See the `retain_rule` block. Max of 1 per action.
	RetainRule Dlm_LifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRule `json:"retainRule,omitempty" yaml:"retainRule,omitempty"`

	// The target Region or the Amazon Resource Name (ARN) of the target Outpost for the snapshot copies.
	Target string `json:"target,omitempty" yaml:"target,omitempty"`
}
