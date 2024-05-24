package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementSizeConstraintStatement struct {
	// Operator to use to compare the request part to the size setting. Valid values include: `EQ`, `NE`, `LE`, `LT`, `GE`, or `GT`.
	ComparisonOperator string `json:"comparisonOperator,omitempty" yaml:"comparisonOperator,omitempty"`

	// Part of a web request that you want AWS WAF to inspect. See `field_to_match` below for details.
	FieldToMatch Wafv2_WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementSizeConstraintStatementFieldToMatch `json:"fieldToMatch,omitempty" yaml:"fieldToMatch,omitempty"`

	// Size, in bytes, to compare to the request part, after any transformations. Valid values are integers between 0 and 21474836480, inclusive.
	Size int `json:"size,omitempty" yaml:"size,omitempty"`

	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection. At least one transformation is required. See `text_transformation` below for details.
	TextTransformations []Wafv2_WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementSizeConstraintStatementTextTransformation `json:"textTransformations,omitempty" yaml:"textTransformations,omitempty"`
}
