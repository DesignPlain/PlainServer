package types

type Route53_getResolverFirewallRulesFirewallRule struct {
	// The custom DNS record to send back in response to the query.
	BlockOverrideDomain string `json:"blockOverrideDomain,omitempty" yaml:"blockOverrideDomain,omitempty"`

	// The recommended amount of time, in seconds, for the DNS resolver or web browser to cache the provided override record.
	BlockOverrideTtl int `json:"blockOverrideTtl,omitempty" yaml:"blockOverrideTtl,omitempty"`

	// The way that you want DNS Firewall to block the request.
	BlockResponse string `json:"blockResponse,omitempty" yaml:"blockResponse,omitempty"`

	// A unique string defined by you to identify the request.
	CreatorRequestId string `json:"creatorRequestId,omitempty" yaml:"creatorRequestId,omitempty"`

	// The ID of the domain list that's used in the rule.
	FirewallDomainListId string `json:"firewallDomainListId,omitempty" yaml:"firewallDomainListId,omitempty"`

	// The unique identifier of the firewall rule group that you want to retrieve the rules for.
	FirewallRuleGroupId string `json:"firewallRuleGroupId,omitempty" yaml:"firewallRuleGroupId,omitempty"`

	// The date and time that the rule was last modified, in Unix time format and Coordinated Universal Time (UTC).
	ModificationTime string `json:"modificationTime,omitempty" yaml:"modificationTime,omitempty"`

	// The action that DNS Firewall should take on a DNS query when it matches one of the domains in the rule's domain list.
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	// The name of the rule.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The date and time that the rule was created, in Unix time format and Coordinated Universal Time (UTC).
	CreationTime string `json:"creationTime,omitempty" yaml:"creationTime,omitempty"`

	// The setting that determines the processing order of the rules in a rule group.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// The DNS record's type.
	BlockOverrideDnsType string `json:"blockOverrideDnsType,omitempty" yaml:"blockOverrideDnsType,omitempty"`
}
