package types

type Securityposture_PosturePolicySet struct {
	// Description of the policy set.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   List of security policy
	   Structure is documented below.
	*/
	Policies []Securityposture_PosturePolicySetPolicy `json:"policies,omitempty" yaml:"policies,omitempty"`

	// ID of the policy set.
	PolicySetId string `json:"policySetId,omitempty" yaml:"policySetId,omitempty"`
}
