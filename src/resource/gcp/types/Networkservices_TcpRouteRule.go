package types

type Networkservices_TcpRouteRule struct {
	/*
	   A detailed rule defining how to route traffic.
	   Structure is documented below.
	*/
	Action Networkservices_TcpRouteRuleAction `json:"action,omitempty" yaml:"action,omitempty"`

	/*
	   RouteMatch defines the predicate used to match requests to a given action. Multiple match types are "OR"ed for evaluation.
	   If no routeMatch field is specified, this rule will unconditionally match traffic.
	   Structure is documented below.
	*/
	Matches []Networkservices_TcpRouteRuleMatch `json:"matches,omitempty" yaml:"matches,omitempty"`
}
