package types

type Wafv2_RuleGroupRuleStatementRateBasedStatementCustomKeyCookie struct {
	// A friendly name of the rule group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection. They are used in rate-based rule statements, to transform request components before using them as custom aggregation keys. Atleast one transformation is required. See Text Transformation above for details.
	TextTransformations []Wafv2_RuleGroupRuleStatementRateBasedStatementCustomKeyCookieTextTransformation `json:"textTransformations,omitempty" yaml:"textTransformations,omitempty"`
}
