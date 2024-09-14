package types

type Networkfirewall_RuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRuleRuleDefinitionMatchAttributes struct {
	// Set of configuration blocks describing the destination ports to inspect for. If not specified, this matches with any destination port. See Destination Port below for details.
	DestinationPorts []Networkfirewall_RuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRuleRuleDefinitionMatchAttributesDestinationPort `json:"destinationPorts,omitempty" yaml:"destinationPorts,omitempty"`

	// Set of configuration blocks describing the destination IP address and address ranges to inspect for, in CIDR notation. If not specified, this matches with any destination address. See Destination below for details.
	Destinations []Networkfirewall_RuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRuleRuleDefinitionMatchAttributesDestination `json:"destinations,omitempty" yaml:"destinations,omitempty"`

	// Set of protocols to inspect for, specified using the protocol's assigned internet protocol number (IANA). If not specified, this matches with any protocol.
	Protocols []int `json:"protocols,omitempty" yaml:"protocols,omitempty"`

	// Set of configuration blocks describing the source ports to inspect for. If not specified, this matches with any source port. See Source Port below for details.
	SourcePorts []Networkfirewall_RuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRuleRuleDefinitionMatchAttributesSourcePort `json:"sourcePorts,omitempty" yaml:"sourcePorts,omitempty"`

	// Set of configuration blocks describing the source IP address and address ranges to inspect for, in CIDR notation. If not specified, this matches with any source address. See Source below for details.
	Sources []Networkfirewall_RuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRuleRuleDefinitionMatchAttributesSource `json:"sources,omitempty" yaml:"sources,omitempty"`

	// Set of configuration blocks containing the TCP flags and masks to inspect for. If not specified, this matches with any settings.
	TcpFlags []Networkfirewall_RuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRuleRuleDefinitionMatchAttributesTcpFlag `json:"tcpFlags,omitempty" yaml:"tcpFlags,omitempty"`
}
