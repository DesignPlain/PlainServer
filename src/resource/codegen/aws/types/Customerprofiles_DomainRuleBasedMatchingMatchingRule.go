package types

type Customerprofiles_DomainRuleBasedMatchingMatchingRule struct {
	// A single rule level of the `match_rules`. Configures how the rule-based matching process should match profiles.
	Rules []string `json:"rules,omitempty" yaml:"rules,omitempty"`
}
