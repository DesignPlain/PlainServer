package types

type Wafv2_WebAclRuleStatementRateBasedStatementScopeDownStatementIpSetReferenceStatementIpSetForwardedIpConfig struct {
	// Match status to assign to the web request if the request doesn't have a valid IP address in the specified position. Valid values include: `MATCH` or `NO_MATCH`.
	FallbackBehavior string `json:"fallbackBehavior,omitempty" yaml:"fallbackBehavior,omitempty"`

	// Name of the HTTP header to use for the IP address.
	HeaderName string `json:"headerName,omitempty" yaml:"headerName,omitempty"`

	// Position in the header to search for the IP address. Valid values include: `FIRST`, `LAST`, or `ANY`. If `ANY` is specified and the header contains more than 10 IP addresses, AWS WAFv2 inspects the last 10.
	Position string `json:"position,omitempty" yaml:"position,omitempty"`
}
