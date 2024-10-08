package types

type Networksecurity_AuthorizationPolicyRule struct {
	/*
	   List of attributes for the traffic source. All of the sources must match. A source is a match if both principals and ipBlocks match.
	   If not set, the action specified in the 'action' field will be applied without any rule checks for the source.
	   Structure is documented below.
	*/
	Sources []Networksecurity_AuthorizationPolicyRuleSource `json:"sources,omitempty" yaml:"sources,omitempty"`

	/*
	   List of attributes for the traffic destination. All of the destinations must match. A destination is a match if a request matches all the specified hosts, ports, methods and headers.
	   If not set, the action specified in the 'action' field will be applied without any rule checks for the destination.
	   Structure is documented below.
	*/
	Destinations []Networksecurity_AuthorizationPolicyRuleDestination `json:"destinations,omitempty" yaml:"destinations,omitempty"`
}
