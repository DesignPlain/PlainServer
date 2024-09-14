package types

type Networkservices_GrpcRouteRuleActionFaultInjectionPolicyDelay struct {
	// Specify a fixed delay before forwarding the request.
	FixedDelay string `json:"fixedDelay,omitempty" yaml:"fixedDelay,omitempty"`

	// The percentage of traffic on which delay will be injected.
	Percentage int `json:"percentage,omitempty" yaml:"percentage,omitempty"`
}
