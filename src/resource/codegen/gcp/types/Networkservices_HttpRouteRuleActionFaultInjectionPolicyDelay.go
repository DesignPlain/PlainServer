package types

type Networkservices_HttpRouteRuleActionFaultInjectionPolicyDelay struct {
	// The percentage of traffic on which delay will be injected.
	Percentage int `json:"percentage,omitempty" yaml:"percentage,omitempty"`

	// Specify a fixed delay before forwarding the request.
	FixedDelay string `json:"fixedDelay,omitempty" yaml:"fixedDelay,omitempty"`
}
