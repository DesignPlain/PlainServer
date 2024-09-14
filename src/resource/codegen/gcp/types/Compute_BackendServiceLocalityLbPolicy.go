package types

type Compute_BackendServiceLocalityLbPolicy struct {
	/*
	   The configuration for a custom policy implemented by the user and
	   deployed with the client.
	   Structure is documented below.
	*/
	CustomPolicy Compute_BackendServiceLocalityLbPolicyCustomPolicy `json:"customPolicy,omitempty" yaml:"customPolicy,omitempty"`

	/*
	   The configuration for a built-in load balancing policy.
	   Structure is documented below.
	*/
	Policy Compute_BackendServiceLocalityLbPolicyPolicy `json:"policy,omitempty" yaml:"policy,omitempty"`
}
