package types

type Cloudrun_getServiceTemplateSpecContainerPort struct {
	// The name of the Cloud Run Service.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Protocol for port. Must be "TCP". Defaults to "TCP".
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	// Port number the container listens on. This must be a valid port number (between 1 and 65535). Defaults to "8080".
	ContainerPort int `json:"containerPort,omitempty" yaml:"containerPort,omitempty"`
}
