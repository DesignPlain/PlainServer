package route53

type ResolverFirewallRuleGroupAssociation struct {
	// The unique identifier of the firewall rule group.
	FirewallRuleGroupId string `json:"firewallRuleGroupId,omitempty" yaml:"firewallRuleGroupId,omitempty"`

	// If enabled, this setting disallows modification or removal of the association, to help prevent against accidentally altering DNS firewall protections. Valid values: `ENABLED`, `DISABLED`.
	MutationProtection string `json:"mutationProtection,omitempty" yaml:"mutationProtection,omitempty"`

	// A name that lets you identify the rule group association, to manage and use it.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The setting that determines the processing order of the rule group among the rule groups that you associate with the specified VPC. DNS Firewall filters VPC traffic starting from the rule group with the lowest numeric priority setting.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The unique identifier of the VPC that you want to associate with the rule group.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
