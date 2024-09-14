package types

type Compute_URLMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy struct {
	/*
	   The specification for how client requests are aborted as part of fault injection.
	   Structure is documented below.
	*/
	Abort Compute_URLMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort `json:"abort,omitempty" yaml:"abort,omitempty"`

	/*
	   The specification for how client requests are delayed as part of fault injection, before being sent to a backend service.
	   Structure is documented below.
	*/
	Delay Compute_URLMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay `json:"delay,omitempty" yaml:"delay,omitempty"`
}
