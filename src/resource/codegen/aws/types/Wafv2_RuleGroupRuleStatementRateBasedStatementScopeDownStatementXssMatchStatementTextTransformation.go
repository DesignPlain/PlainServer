package types

type Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementXssMatchStatementTextTransformation struct {
	// The transformation to apply, please refer to the Text Transformation [documentation](https://docs.aws.amazon.com/waf/latest/APIReference/API_TextTransformation.html) for more details.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The relative processing order for multiple transformations that are defined for a rule statement. AWS WAF processes all transformations, from lowest priority to highest, before inspecting the transformed content.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`
}
