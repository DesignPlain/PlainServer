package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementGeoMatchStatementForwardedIpConfig struct {
	// Name of the HTTP header to use for the IP address.
	HeaderName string `json:"headerName,omitempty" yaml:"headerName,omitempty"`

	// Match status to assign to the web request if the request doesn't have a valid IP address in the specified position. Valid values include: `MATCH` or `NO_MATCH`.
	FallbackBehavior string `json:"fallbackBehavior,omitempty" yaml:"fallbackBehavior,omitempty"`
}
