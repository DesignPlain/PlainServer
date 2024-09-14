package types

type Networkservices_GrpcRouteRuleActionFaultInjectionPolicy struct {
	/*
	   Specification of how client requests are aborted as part of fault injection before being sent to a destination.
	   Structure is documented below.
	*/
	Abort Networkservices_GrpcRouteRuleActionFaultInjectionPolicyAbort `json:"abort,omitempty" yaml:"abort,omitempty"`

	/*
	   Specification of how client requests are delayed as part of fault injection before being sent to a destination.
	   Structure is documented below.
	*/
	Delay Networkservices_GrpcRouteRuleActionFaultInjectionPolicyDelay `json:"delay,omitempty" yaml:"delay,omitempty"`
}
