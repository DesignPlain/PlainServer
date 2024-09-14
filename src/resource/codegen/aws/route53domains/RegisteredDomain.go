package route53domains

import types "libds/aws/types"

type RegisteredDomain struct {
	// Details about the domain administrative contact. See Contact Blocks for more details.
	AdminContact types.Route53domains_RegisteredDomainAdminContact `json:"adminContact,omitempty" yaml:"adminContact,omitempty"`

	// Whether the domain registration is set to renew automatically. Default: `true`.
	AutoRenew bool `json:"autoRenew,omitempty" yaml:"autoRenew,omitempty"`

	// Whether domain billing contact information is concealed from WHOIS queries. Default: `true`.
	BillingPrivacy bool `json:"billingPrivacy,omitempty" yaml:"billingPrivacy,omitempty"`

	// The list of nameservers for the domain. See `name_server` Blocks for more details.
	NameServers []types.Route53domains_RegisteredDomainNameServer `json:"nameServers,omitempty" yaml:"nameServers,omitempty"`

	// Details about the domain registrant. See Contact Blocks for more details.
	RegistrantContact types.Route53domains_RegisteredDomainRegistrantContact `json:"registrantContact,omitempty" yaml:"registrantContact,omitempty"`

	// Whether domain technical contact information is concealed from WHOIS queries. Default: `true`.
	TechPrivacy bool `json:"techPrivacy,omitempty" yaml:"techPrivacy,omitempty"`

	// Whether domain administrative contact information is concealed from WHOIS queries. Default: `true`.
	AdminPrivacy bool `json:"adminPrivacy,omitempty" yaml:"adminPrivacy,omitempty"`

	// Details about the domain billing contact. See Contact Blocks for more details.
	BillingContact types.Route53domains_RegisteredDomainBillingContact `json:"billingContact,omitempty" yaml:"billingContact,omitempty"`

	// The name of the registered domain.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// Whether domain registrant contact information is concealed from WHOIS queries. Default: `true`.
	RegistrantPrivacy bool `json:"registrantPrivacy,omitempty" yaml:"registrantPrivacy,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Details about the domain technical contact. See Contact Blocks for more details.
	TechContact types.Route53domains_RegisteredDomainTechContact `json:"techContact,omitempty" yaml:"techContact,omitempty"`

	// Whether the domain is locked for transfer. Default: `true`.
	TransferLock bool `json:"transferLock,omitempty" yaml:"transferLock,omitempty"`
}
