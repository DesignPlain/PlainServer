package types

type Wafv2_RuleGroupRuleStatementRegexPatternSetReferenceStatement struct {
	// The Amazon Resource Name (ARN) of the Regex Pattern Set that this statement references.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// The part of a web request that you want AWS WAF to inspect. See Field to Match below for details.
	FieldToMatch Wafv2_RuleGroupRuleStatementRegexPatternSetReferenceStatementFieldToMatch `json:"fieldToMatch,omitempty" yaml:"fieldToMatch,omitempty"`

	/*
	   Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	   At least one required.
	   See Text Transformation below for details.
	*/
	TextTransformations []Wafv2_RuleGroupRuleStatementRegexPatternSetReferenceStatementTextTransformation `json:"textTransformations,omitempty" yaml:"textTransformations,omitempty"`
}
