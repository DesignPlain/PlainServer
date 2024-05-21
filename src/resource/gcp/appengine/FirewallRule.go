package appengine

type FirewallRule struct {
	// An optional string description of this rule.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   A positive integer that defines the order of rule evaluation.
	   Rules with the lowest priority are evaluated first.
	   A default rule at priority Int32.MaxValue matches all IPv4 and
	   IPv6 traffic when no previous rule matches. Only the action of
	   this rule can be modified by the user.
	*/
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// IP address or range, defined using CIDR notation, of requests that this rule applies to.
	SourceRange string `json:"sourceRange,omitempty" yaml:"sourceRange,omitempty"`

	/*
	   The action to take if this rule matches.
	   Possible values are: `UNSPECIFIED_ACTION`, `ALLOW`, `DENY`.


	   - - -
	*/
	Action string `json:"action,omitempty" yaml:"action,omitempty"`
}
