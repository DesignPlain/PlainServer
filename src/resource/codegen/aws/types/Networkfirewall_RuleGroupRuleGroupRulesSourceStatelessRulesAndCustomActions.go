package types

type Networkfirewall_RuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActions struct {
	// Set of configuration blocks containing custom action definitions that are available for use by the set of `stateless rule`. See Custom Action below for details.
	CustomActions []Networkfirewall_RuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsCustomAction `json:"customActions,omitempty" yaml:"customActions,omitempty"`

	// Set of configuration blocks containing the stateless rules for use in the stateless rule group. See Stateless Rule below for details.
	StatelessRules []Networkfirewall_RuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRule `json:"statelessRules,omitempty" yaml:"statelessRules,omitempty"`
}
