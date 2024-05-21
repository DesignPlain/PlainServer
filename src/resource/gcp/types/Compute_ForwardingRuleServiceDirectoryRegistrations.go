package types

type Compute_ForwardingRuleServiceDirectoryRegistrations struct {
	// Service Directory namespace to register the forwarding rule under.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`

	// Service Directory service to register the forwarding rule under.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}
