package types

type Ecs_ServiceServiceConnectConfigurationService struct {
	// The name of one of the `portMappings` from all the containers in the task definition of this Amazon ECS service.
	PortName string `json:"portName,omitempty" yaml:"portName,omitempty"`

	// Configuration timeouts for Service Connect
	Timeout Ecs_ServiceServiceConnectConfigurationServiceTimeout `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	// The configuration for enabling Transport Layer Security (TLS)
	Tls Ecs_ServiceServiceConnectConfigurationServiceTls `json:"tls,omitempty" yaml:"tls,omitempty"`

	// The list of client aliases for this Service Connect service. You use these to assign names that can be used by client applications. The maximum number of client aliases that you can have in this list is 1. See below.
	ClientAlias []Ecs_ServiceServiceConnectConfigurationServiceClientAlias `json:"clientAlias,omitempty" yaml:"clientAlias,omitempty"`

	// The name of the new AWS Cloud Map service that Amazon ECS creates for this Amazon ECS service.
	DiscoveryName string `json:"discoveryName,omitempty" yaml:"discoveryName,omitempty"`

	// The port number for the Service Connect proxy to listen on.
	IngressPortOverride int `json:"ingressPortOverride,omitempty" yaml:"ingressPortOverride,omitempty"`
}
