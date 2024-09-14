package types

type Wafv2_RuleGroupRuleStatementByteMatchStatement struct {
	// The part of a web request that you want AWS WAF to inspect. See Field to Match below for details.
	FieldToMatch Wafv2_RuleGroupRuleStatementByteMatchStatementFieldToMatch `json:"fieldToMatch,omitempty" yaml:"fieldToMatch,omitempty"`

	// The area within the portion of a web request that you want AWS WAF to search for `search_string`. Valid values include the following: `EXACTLY`, `STARTS_WITH`, `ENDS_WITH`, `CONTAINS`, `CONTAINS_WORD`. See the AWS [documentation](https://docs.aws.amazon.com/waf/latest/APIReference/API_ByteMatchStatement.html) for more information.
	PositionalConstraint string `json:"positionalConstraint,omitempty" yaml:"positionalConstraint,omitempty"`

	// A string value that you want AWS WAF to search for. AWS WAF searches only in the part of web requests that you designate for inspection in `field_to_match`. The maximum length of the value is 50 bytes.
	SearchString string `json:"searchString,omitempty" yaml:"searchString,omitempty"`

	/*
	   Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	   At least one required.
	   See Text Transformation below for details.
	*/
	TextTransformations []Wafv2_RuleGroupRuleStatementByteMatchStatementTextTransformation `json:"textTransformations,omitempty" yaml:"textTransformations,omitempty"`
}
