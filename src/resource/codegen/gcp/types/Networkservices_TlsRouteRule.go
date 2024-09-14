package types

type Networkservices_TlsRouteRule struct {
	/*
	   Required. A detailed rule defining how to route traffic.
	   Structure is documented below.
	*/
	Action Networkservices_TlsRouteRuleAction `json:"action,omitempty" yaml:"action,omitempty"`

	/*
	   Matches define the predicate used to match requests to a given action.
	   Structure is documented below.
	*/
	Matches []Networkservices_TlsRouteRuleMatch `json:"matches,omitempty" yaml:"matches,omitempty"`
}
