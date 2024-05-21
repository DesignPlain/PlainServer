package types

type Securityposture_PosturePolicySetPolicyComplianceStandard struct {
	// Mapping of security controls for the policy.
	Control string `json:"control,omitempty" yaml:"control,omitempty"`

	// Mapping of compliance standards for the policy.
	Standard string `json:"standard,omitempty" yaml:"standard,omitempty"`
}
