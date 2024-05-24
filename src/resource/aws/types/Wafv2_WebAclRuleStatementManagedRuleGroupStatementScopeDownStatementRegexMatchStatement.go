package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatement struct {
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection. At least one transformation is required. See `text_transformation` below for details.
	TextTransformations []Wafv2_WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementTextTransformation `json:"textTransformations,omitempty" yaml:"textTransformations,omitempty"`

	// The part of a web request that you want AWS WAF to inspect. See `field_to_match` below for details.
	FieldToMatch Wafv2_WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatch `json:"fieldToMatch,omitempty" yaml:"fieldToMatch,omitempty"`

	// String representing the regular expression. Minimum of `1` and maximum of `512` characters.
	RegexString string `json:"regexString,omitempty" yaml:"regexString,omitempty"`
}
