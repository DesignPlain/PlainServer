package route53

type ResolverFirewallDomainList struct {
	// A map of tags to assign to the resource. f configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A array of domains for the firewall domain list.
	Domains []string `json:"domains,omitempty" yaml:"domains,omitempty"`

	// A name that lets you identify the domain list, to manage and use it.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
