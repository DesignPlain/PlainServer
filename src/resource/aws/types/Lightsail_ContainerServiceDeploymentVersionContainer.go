package types

type Lightsail_ContainerServiceDeploymentVersionContainer struct {
	// The launch command for the container. A list of string.
	Commands []string `json:"commands,omitempty" yaml:"commands,omitempty"`

	// The name for the container.
	ContainerName string `json:"containerName,omitempty" yaml:"containerName,omitempty"`

	// A key-value map of the environment variables of the container.
	Environment map[string]string `json:"environment,omitempty" yaml:"environment,omitempty"`

	// The name of the image used for the container. Container images sourced from your Lightsail container service, that are registered and stored on your service, start with a colon (`:`). For example, `:container-service-1.mystaticwebsite.1`. Container images sourced from a public registry like Docker Hub don't start with a colon. For example, `nginx:latest` or `nginx`.
	Image string `json:"image,omitempty" yaml:"image,omitempty"`

	// A key-value map of the open firewall ports of the container. Valid values: `HTTP`, `HTTPS`, `TCP`, `UDP`.
	Ports map[string]string `json:"ports,omitempty" yaml:"ports,omitempty"`
}
