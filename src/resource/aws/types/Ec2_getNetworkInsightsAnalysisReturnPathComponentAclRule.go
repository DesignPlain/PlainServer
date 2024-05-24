package types

type Ec2_getNetworkInsightsAnalysisReturnPathComponentAclRule struct {
	//
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	//
	RuleAction string `json:"ruleAction,omitempty" yaml:"ruleAction,omitempty"`

	//
	RuleNumber int `json:"ruleNumber,omitempty" yaml:"ruleNumber,omitempty"`

	//
	Cidr string `json:"cidr,omitempty" yaml:"cidr,omitempty"`

	//
	Egress bool `json:"egress,omitempty" yaml:"egress,omitempty"`

	//
	PortRanges []Ec2_getNetworkInsightsAnalysisReturnPathComponentAclRulePortRange `json:"portRanges,omitempty" yaml:"portRanges,omitempty"`
}
