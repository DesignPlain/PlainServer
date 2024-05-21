package types

type Securityposture_PosturePolicySetPolicyConstraintOrgPolicyConstraintPolicyRuleValues struct {
	// List of values allowed at this resource.
	AllowedValues []string `json:"allowedValues,omitempty" yaml:"allowedValues,omitempty"`

	// List of values denied at this resource.
	DeniedValues []string `json:"deniedValues,omitempty" yaml:"deniedValues,omitempty"`
}
