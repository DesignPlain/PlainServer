package types

type Compute_URLMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay struct {
	/*
	   Specifies the value of the fixed delay interval.
	   Structure is documented below.
	*/
	FixedDelay Compute_URLMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay `json:"fixedDelay,omitempty" yaml:"fixedDelay,omitempty"`

	/*
	   The percentage of traffic (connections/operations/requests) on which delay will be introduced as part of fault injection.
	   The value must be between 0.0 and 100.0 inclusive.
	*/
	Percentage float64 `json:"percentage,omitempty" yaml:"percentage,omitempty"`
}
