package types

type Networkservices_TcpRouteRuleAction struct {
	/*
	   The destination services to which traffic should be forwarded. At least one destination service is required.
	   Structure is documented below.
	*/
	Destinations []Networkservices_TcpRouteRuleActionDestination `json:"destinations,omitempty" yaml:"destinations,omitempty"`

	// If true, Router will use the destination IP and port of the original connection as the destination of the request.
	OriginalDestination bool `json:"originalDestination,omitempty" yaml:"originalDestination,omitempty"`
}
