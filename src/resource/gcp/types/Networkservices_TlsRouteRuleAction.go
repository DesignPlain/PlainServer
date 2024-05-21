package types

type Networkservices_TlsRouteRuleAction struct {
	/*
	   The destination to which traffic should be forwarded.
	   Structure is documented below.
	*/
	Destinations []Networkservices_TlsRouteRuleActionDestination `json:"destinations,omitempty" yaml:"destinations,omitempty"`
}
