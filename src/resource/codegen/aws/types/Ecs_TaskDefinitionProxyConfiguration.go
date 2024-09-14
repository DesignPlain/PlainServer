package types

type Ecs_TaskDefinitionProxyConfiguration struct {
	// Set of network configuration parameters to provide the Container Network Interface (CNI) plugin, specified a key-value mapping.
	Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`

	// Proxy type. The default value is `APPMESH`. The only supported value is `APPMESH`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Name of the container that will serve as the App Mesh proxy.
	ContainerName string `json:"containerName,omitempty" yaml:"containerName,omitempty"`
}
