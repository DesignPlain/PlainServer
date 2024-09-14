package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementByteMatchStatement struct {
	// Part of a web request that you want AWS WAF to inspect. See `field_to_match` below for details.
	FieldToMatch Wafv2_WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatch `json:"fieldToMatch,omitempty" yaml:"fieldToMatch,omitempty"`

	// Area within the portion of a web request that you want AWS WAF to search for `search_string`. Valid values include the following: `EXACTLY`, `STARTS_WITH`, `ENDS_WITH`, `CONTAINS`, `CONTAINS_WORD`. See the AWS [documentation](https://docs.aws.amazon.com/waf/latest/APIReference/API_ByteMatchStatement.html) for more information.
	PositionalConstraint string `json:"positionalConstraint,omitempty" yaml:"positionalConstraint,omitempty"`

	// String value that you want AWS WAF to search for. AWS WAF searches only in the part of web requests that you designate for inspection in `field_to_match`. The maximum length of the value is 50 bytes.
	SearchString string `json:"searchString,omitempty" yaml:"searchString,omitempty"`

	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection. At least one transformation is required. See `text_transformation` below for details.
	TextTransformations []Wafv2_WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementTextTransformation `json:"textTransformations,omitempty" yaml:"textTransformations,omitempty"`
}
