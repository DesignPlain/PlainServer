package types

type Securityposture_PosturePolicySetPolicy struct {
	/*
	   Mapping for policy to security standards and controls.
	   Structure is documented below.
	*/
	ComplianceStandards []Securityposture_PosturePolicySetPolicyComplianceStandard `json:"complianceStandards,omitempty" yaml:"complianceStandards,omitempty"`

	/*
	   Policy constraint definition.It can have the definition of one of following constraints: orgPolicyConstraint orgPolicyConstraintCustom securityHealthAnalyticsModule securityHealthAnalyticsCustomModule
	   Structure is documented below.
	*/
	Constraint Securityposture_PosturePolicySetPolicyConstraint `json:"constraint,omitempty" yaml:"constraint,omitempty"`

	// Description of the policy.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// ID of the policy.
	PolicyId string `json:"policyId,omitempty" yaml:"policyId,omitempty"`
}
