package types

type Compute_getForwardingRuleServiceDirectoryRegistration struct {
	// Service Directory service to register the forwarding rule under.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`

	// Service Directory namespace to register the forwarding rule under.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
}
