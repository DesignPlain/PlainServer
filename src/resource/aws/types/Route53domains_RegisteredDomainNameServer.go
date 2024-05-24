package types

type Route53domains_RegisteredDomainNameServer struct {
	// The fully qualified host name of the name server.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Glue IP addresses of a name server. The list can contain only one IPv4 and one IPv6 address.
	GlueIps []string `json:"glueIps,omitempty" yaml:"glueIps,omitempty"`
}
