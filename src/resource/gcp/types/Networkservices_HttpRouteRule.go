package types

type Networkservices_HttpRouteRule struct {
	/*
	   The detailed rule defining how to route matched traffic.
	   Structure is documented below.
	*/
	Action Networkservices_HttpRouteRuleAction `json:"action,omitempty" yaml:"action,omitempty"`

	/*
	   A list of matches define conditions used for matching the rule against incoming HTTP requests. Each match is independent, i.e. this rule will be matched if ANY one of the matches is satisfied.
	   If no matches field is specified, this rule will unconditionally match traffic.
	   If a default rule is desired to be configured, add a rule with no matches specified to the end of the rules list.
	   Structure is documented below.
	*/
	Matches []Networkservices_HttpRouteRuleMatch `json:"matches,omitempty" yaml:"matches,omitempty"`
}
