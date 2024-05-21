package types



type Securityposture_PosturePolicySetPolicyConstraintOrgPolicyConstraintCustom struct {
	/*
	   Definition of policy rules
	   Structure is documented below.
	*/
	PolicyRules []Securityposture_PosturePolicySetPolicyConstraintOrgPolicyConstraintCustomPolicyRule `json:"policyRules,omitempty" yaml:"policyRules,omitempty"`

	/*
	   Organization policy custom constraint definition.
	   Structure is documented below.
	*/
	CustomConstraint Securityposture_PosturePolicySetPolicyConstraintOrgPolicyConstraintCustomCustomConstraint `json:"customConstraint,omitempty" yaml:"customConstraint,omitempty"`
}
