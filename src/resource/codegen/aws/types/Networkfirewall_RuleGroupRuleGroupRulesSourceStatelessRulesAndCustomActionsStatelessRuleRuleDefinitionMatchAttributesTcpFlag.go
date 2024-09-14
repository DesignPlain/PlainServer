package types

type Networkfirewall_RuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRuleRuleDefinitionMatchAttributesTcpFlag struct {
	/*
	   Set of flags to look for in a packet. This setting can only specify values that are also specified in `masks`.
	   Valid values: `FIN`, `SYN`, `RST`, `PSH`, `ACK`, `URG`, `ECE`, `CWR`.
	*/
	Flags []string `json:"flags,omitempty" yaml:"flags,omitempty"`

	/*
	   Set of flags to consider in the inspection. To inspect all flags, leave this empty.
	   Valid values: `FIN`, `SYN`, `RST`, `PSH`, `ACK`, `URG`, `ECE`, `CWR`.
	*/
	Masks []string `json:"masks,omitempty" yaml:"masks,omitempty"`
}
