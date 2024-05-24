package route53

import types "DesignSphere_Server/src/resource/aws/types"

type ResolverRule struct {
	/*
	   Configuration block(s) indicating the IPs that you want Resolver to forward DNS queries to (documented below).
	   This argument should only be specified for `FORWARD` type rules.
	*/
	TargetIps []types.Route53_ResolverRuleTargetIp `json:"targetIps,omitempty" yaml:"targetIps,omitempty"`

	// DNS queries for this domain name are forwarded to the IP addresses that are specified using `target_ip`.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// A friendly name that lets you easily find a rule in the Resolver dashboard in the Route 53 console.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the outbound resolver endpoint that you want to use to route DNS queries to the IP addresses that you specify using `target_ip`.
	   This argument should only be specified for `FORWARD` type rules.
	*/
	ResolverEndpointId string `json:"resolverEndpointId,omitempty" yaml:"resolverEndpointId,omitempty"`

	// The rule type. Valid values are `FORWARD`, `SYSTEM` and `RECURSIVE`.
	RuleType string `json:"ruleType,omitempty" yaml:"ruleType,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
