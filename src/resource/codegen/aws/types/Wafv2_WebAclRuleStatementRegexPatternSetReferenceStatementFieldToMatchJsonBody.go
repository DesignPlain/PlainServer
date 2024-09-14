package types

type Wafv2_WebAclRuleStatementRegexPatternSetReferenceStatementFieldToMatchJsonBody struct {
	// What to do when JSON parsing fails. Defaults to evaluating up to the first parsing failure. Valid values are `EVALUATE_AS_STRING`, `MATCH` and `NO_MATCH`.
	InvalidFallbackBehavior string `json:"invalidFallbackBehavior,omitempty" yaml:"invalidFallbackBehavior,omitempty"`

	// The patterns to look for in the JSON body. You must specify exactly one setting: either `all` or `included_paths`. See [JsonMatchPattern](https://docs.aws.amazon.com/waf/latest/APIReference/API_JsonMatchPattern.html) for details.
	MatchPattern Wafv2_WebAclRuleStatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPattern `json:"matchPattern,omitempty" yaml:"matchPattern,omitempty"`

	// The parts of the JSON to match against using the `match_pattern`. Valid values are `ALL`, `KEY` and `VALUE`.
	MatchScope string `json:"matchScope,omitempty" yaml:"matchScope,omitempty"`

	// What to do if the body is larger than can be inspected. Valid values are `CONTINUE` (default), `MATCH` and `NO_MATCH`.
	OversizeHandling string `json:"oversizeHandling,omitempty" yaml:"oversizeHandling,omitempty"`
}
