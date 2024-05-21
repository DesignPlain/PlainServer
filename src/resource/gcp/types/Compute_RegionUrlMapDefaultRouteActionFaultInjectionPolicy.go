package types

type Compute_RegionUrlMapDefaultRouteActionFaultInjectionPolicy struct {
	/*
	   The specification for how client requests are aborted as part of fault injection.
	   Structure is documented below.
	*/
	Abort Compute_RegionUrlMapDefaultRouteActionFaultInjectionPolicyAbort `json:"abort,omitempty" yaml:"abort,omitempty"`

	/*
	   The specification for how client requests are delayed as part of fault injection, before being sent to a backend service.
	   Structure is documented below.
	*/
	Delay Compute_RegionUrlMapDefaultRouteActionFaultInjectionPolicyDelay `json:"delay,omitempty" yaml:"delay,omitempty"`
}
