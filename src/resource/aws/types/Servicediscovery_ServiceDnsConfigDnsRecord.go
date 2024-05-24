package types

type Servicediscovery_ServiceDnsConfigDnsRecord struct {
	// The amount of time, in seconds, that you want DNS resolvers to cache the settings for this resource record set.
	Ttl int `json:"ttl,omitempty" yaml:"ttl,omitempty"`

	// The type of the resource, which indicates the value that Amazon Route 53 returns in response to DNS queries. Valid Values: A, AAAA, SRV, CNAME
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
