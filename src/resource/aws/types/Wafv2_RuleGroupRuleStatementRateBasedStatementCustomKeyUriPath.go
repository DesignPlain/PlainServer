package types

type Wafv2_RuleGroupRuleStatementRateBasedStatementCustomKeyUriPath struct {
	/*
	   Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	   At least one required.
	   See Text Transformation below for details.
	*/
	TextTransformations []Wafv2_RuleGroupRuleStatementRateBasedStatementCustomKeyUriPathTextTransformation `json:"textTransformations,omitempty" yaml:"textTransformations,omitempty"`
}
