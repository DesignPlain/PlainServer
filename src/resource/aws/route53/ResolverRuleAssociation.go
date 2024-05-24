package route53

type ResolverRuleAssociation struct {
	// A name for the association that you're creating between a resolver rule and a VPC.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The ID of the resolver rule that you want to associate with the VPC.
	ResolverRuleId string `json:"resolverRuleId,omitempty" yaml:"resolverRuleId,omitempty"`

	// The ID of the VPC that you want to associate the resolver rule with.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
