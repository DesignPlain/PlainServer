package types

type Servicediscovery_getServiceDnsConfigDnsRecord struct {
	// Amount of time, in seconds, that you want DNS resolvers to cache the settings for this resource record set.
	Ttl int `json:"ttl,omitempty" yaml:"ttl,omitempty"`

	// The type of health check that you want to create, which indicates how Route 53 determines whether an endpoint is healthy. Valid Values: HTTP, HTTPS, TCP
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
