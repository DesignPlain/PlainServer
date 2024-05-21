package types



type Networkservices_GrpcRouteRule struct {
	/*
	   Required. A detailed rule defining how to route traffic.
	   Structure is documented below.
	*/
	Action Networkservices_GrpcRouteRuleAction `json:"action,omitempty" yaml:"action,omitempty"`

	/*
	   Matches define conditions used for matching the rule against incoming gRPC requests.
	   Structure is documented below.
	*/
	Matches []Networkservices_GrpcRouteRuleMatch `json:"matches,omitempty" yaml:"matches,omitempty"`
}
