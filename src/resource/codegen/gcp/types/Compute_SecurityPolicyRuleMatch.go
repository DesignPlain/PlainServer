package types

type Compute_SecurityPolicyRuleMatch struct {
	/*
	   The configuration options available when specifying `versioned_expr`.
	   This field must be specified if `versioned_expr` is specified and cannot be specified if `versioned_expr` is not specified.
	   Structure is documented below.
	*/
	Config Compute_SecurityPolicyRuleMatchConfig `json:"config,omitempty" yaml:"config,omitempty"`

	/*
	   User defined CEVAL expression. A CEVAL expression is used to specify match criteria
	   such as `origin.ip`, `source.region_code` and `contents` in the request header.
	   Structure is documented below.
	*/
	Expr Compute_SecurityPolicyRuleMatchExpr `json:"expr,omitempty" yaml:"expr,omitempty"`

	/*
	   Predefined rule expression. If this field is specified, `config` must also be specified.
	   Available options:
	*/
	VersionedExpr string `json:"versionedExpr,omitempty" yaml:"versionedExpr,omitempty"`
}
