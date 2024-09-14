package types

type Dlm_LifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRule struct {
	//
	CmkArn string `json:"cmkArn,omitempty" yaml:"cmkArn,omitempty"`

	//
	CopyTags bool `json:"copyTags,omitempty" yaml:"copyTags,omitempty"`

	//
	DeprecateRule Dlm_LifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRule `json:"deprecateRule,omitempty" yaml:"deprecateRule,omitempty"`

	//
	Encrypted bool `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`

	//
	RetainRule Dlm_LifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRule `json:"retainRule,omitempty" yaml:"retainRule,omitempty"`

	//
	Target string `json:"target,omitempty" yaml:"target,omitempty"`
}
