package types



type Compute_OrganizationSecurityPolicyRuleMatch struct {
	/*
	   Preconfigured versioned expression. For organization security policy rules,
	   the only supported type is "FIREWALL".
	   Default value is `FIREWALL`.
	   Possible values are: `FIREWALL`.
	*/
	VersionedExpr string `json:"versionedExpr,omitempty" yaml:"versionedExpr,omitempty"`

	/*
	   The configuration options for matching the rule.
	   Structure is documented below.
	*/
	Config Compute_OrganizationSecurityPolicyRuleMatchConfig `json:"config,omitempty" yaml:"config,omitempty"`

	// A description of the rule.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
