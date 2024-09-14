package servicediscovery

import types "libds/aws/types"

type Service struct {
	// A complex type that contains settings for an optional health check. Only for Public DNS namespaces. See `health_check_config` Block for details.
	HealthCheckConfig types.Servicediscovery_ServiceHealthCheckConfig `json:"healthCheckConfig,omitempty" yaml:"healthCheckConfig,omitempty"`

	// The name of the service.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// If present, specifies that the service instances are only discoverable using the `DiscoverInstances` API operation. No DNS records is registered for the service instances. The only valid value is `HTTP`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// A complex type that contains settings for ECS managed health checks. See `health_check_custom_config` Block for details.
	HealthCheckCustomConfig types.Servicediscovery_ServiceHealthCheckCustomConfig `json:"healthCheckCustomConfig,omitempty" yaml:"healthCheckCustomConfig,omitempty"`

	// The ID of the namespace that you want to use to create the service.
	NamespaceId string `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`

	// A map of tags to assign to the service. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The description of the service.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A complex type that contains information about the resource record sets that you want Amazon Route 53 to create when you register an instance. See `dns_config` Block for details.
	DnsConfig types.Servicediscovery_ServiceDnsConfig `json:"dnsConfig,omitempty" yaml:"dnsConfig,omitempty"`

	// A boolean that indicates all instances should be deleted from the service so that the service can be destroyed without error. These instances are not recoverable. Defaults to `false`.
	ForceDestroy bool `json:"forceDestroy,omitempty" yaml:"forceDestroy,omitempty"`
}
