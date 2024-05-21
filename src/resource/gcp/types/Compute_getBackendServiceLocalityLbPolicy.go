package types

type Compute_getBackendServiceLocalityLbPolicy struct {
	/*
	   The configuration for a custom policy implemented by the user and
	   deployed with the client.
	*/
	CustomPolicies []Compute_getBackendServiceLocalityLbPolicyCustomPolicy `json:"customPolicies,omitempty" yaml:"customPolicies,omitempty"`

	// The configuration for a built-in load balancing policy.
	Policies []Compute_getBackendServiceLocalityLbPolicyPolicy `json:"policies,omitempty" yaml:"policies,omitempty"`
}
