package types

type Networkservices_GrpcRouteRuleAction struct {
	// Specifies the timeout for selected route.
	Timeout string `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	/*
	   The destination to which traffic should be forwarded.
	   Structure is documented below.
	*/
	Destinations []Networkservices_GrpcRouteRuleActionDestination `json:"destinations,omitempty" yaml:"destinations,omitempty"`

	/*
	   The specification for fault injection introduced into traffic to test the resiliency of clients to backend service failure.
	   Structure is documented below.
	*/
	FaultInjectionPolicy Networkservices_GrpcRouteRuleActionFaultInjectionPolicy `json:"faultInjectionPolicy,omitempty" yaml:"faultInjectionPolicy,omitempty"`

	/*
	   Specifies the retry policy associated with this route.
	   Structure is documented below.
	*/
	RetryPolicy Networkservices_GrpcRouteRuleActionRetryPolicy `json:"retryPolicy,omitempty" yaml:"retryPolicy,omitempty"`
}
