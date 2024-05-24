package types

type Servicediscovery_ServiceDnsConfig struct {
	// An array that contains one DnsRecord object for each resource record set.
	DnsRecords []Servicediscovery_ServiceDnsConfigDnsRecord `json:"dnsRecords,omitempty" yaml:"dnsRecords,omitempty"`

	// The ID of the namespace to use for DNS configuration.
	NamespaceId string `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`

	// The routing policy that you want to apply to all records that Route 53 creates when you register an instance and specify the service. Valid Values: MULTIVALUE, WEIGHTED
	RoutingPolicy string `json:"routingPolicy,omitempty" yaml:"routingPolicy,omitempty"`
}
