package types

type Servicediscovery_getServiceDnsConfig struct {
	// An array that contains one DnsRecord object for each resource record set. See `dns_records` Block for details.
	DnsRecords []Servicediscovery_getServiceDnsConfigDnsRecord `json:"dnsRecords,omitempty" yaml:"dnsRecords,omitempty"`

	// ID of the namespace that the service belongs to.
	NamespaceId string `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`

	// Routing policy that you want to apply to all records that Route 53 creates when you register an instance and specify the service. Valid Values: MULTIVALUE, WEIGHTED
	RoutingPolicy string `json:"routingPolicy,omitempty" yaml:"routingPolicy,omitempty"`
}
