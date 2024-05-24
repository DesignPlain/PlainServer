package types

type Networkfirewall_RuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRuleRuleDefinitionMatchAttributesSource struct {
	// An IP address or a block of IP addresses in CIDR notation. AWS Network Firewall supports all address ranges for IPv4.
	AddressDefinition string `json:"addressDefinition,omitempty" yaml:"addressDefinition,omitempty"`
}
