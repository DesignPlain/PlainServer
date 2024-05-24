package types

type Ec2_NetworkInsightsAnalysisExplanationAclRule struct {
	//
	PortRanges []Ec2_NetworkInsightsAnalysisExplanationAclRulePortRange `json:"portRanges,omitempty" yaml:"portRanges,omitempty"`

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
}
