package types

type Networkservices_GrpcRouteRuleMatch struct {
	/*
	   Specifies a list of HTTP request headers to match against.
	   Structure is documented below.
	*/
	Headers []Networkservices_GrpcRouteRuleMatchHeader `json:"headers,omitempty" yaml:"headers,omitempty"`

	/*
	   A gRPC method to match against. If this field is empty or omitted, will match all methods.
	   Structure is documented below.
	*/
	Method Networkservices_GrpcRouteRuleMatchMethod `json:"method,omitempty" yaml:"method,omitempty"`
}
