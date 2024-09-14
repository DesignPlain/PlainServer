package types

type Securityposture_PosturePolicySetPolicyConstraintOrgPolicyConstraint struct {
	// Organization policy canned constraint Id
	CannedConstraintId string `json:"cannedConstraintId,omitempty" yaml:"cannedConstraintId,omitempty"`

	/*
	   Definition of policy rules
	   Structure is documented below.
	*/
	PolicyRules []Securityposture_PosturePolicySetPolicyConstraintOrgPolicyConstraintPolicyRule `json:"policyRules,omitempty" yaml:"policyRules,omitempty"`
}
