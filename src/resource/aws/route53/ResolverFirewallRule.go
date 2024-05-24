package route53

type ResolverFirewallRule struct {
	// The action that DNS Firewall should take on a DNS query when it matches one of the domains in the rule's domain list. Valid values: `ALLOW`, `BLOCK`, `ALERT`.
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	// The custom DNS record to send back in response to the query.
	BlockOverrideDomain string `json:"blockOverrideDomain,omitempty" yaml:"blockOverrideDomain,omitempty"`

	// The ID of the domain list that you want to use in the rule.
	FirewallDomainListId string `json:"firewallDomainListId,omitempty" yaml:"firewallDomainListId,omitempty"`

	// The unique identifier of the firewall rule group where you want to create the rule.
	FirewallRuleGroupId string `json:"firewallRuleGroupId,omitempty" yaml:"firewallRuleGroupId,omitempty"`

	// A name that lets you identify the rule, to manage and use it.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The setting that determines the processing order of the rule in the rule group. DNS Firewall processes the rules in a rule group by order of priority, starting from the lowest setting.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// The DNS record's type. This determines the format of the record value that you provided in BlockOverrideDomain. Value values: `CNAME`.
	BlockOverrideDnsType string `json:"blockOverrideDnsType,omitempty" yaml:"blockOverrideDnsType,omitempty"`

	// The recommended amount of time, in seconds, for the DNS resolver or web browser to cache the provided override record. Minimum value of 0. Maximum value of 604800.
	BlockOverrideTtl int `json:"blockOverrideTtl,omitempty" yaml:"blockOverrideTtl,omitempty"`

	// The way that you want DNS Firewall to block the request. Valid values: `NODATA`, `NXDOMAIN`, `OVERRIDE`.
	BlockResponse string `json:"blockResponse,omitempty" yaml:"blockResponse,omitempty"`
}
