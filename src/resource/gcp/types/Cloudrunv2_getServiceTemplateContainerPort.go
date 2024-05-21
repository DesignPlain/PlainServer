package types

type Cloudrunv2_getServiceTemplateContainerPort struct {
	// Port number the container listens on. This must be a valid TCP port number, 0 < containerPort < 65536.
	ContainerPort int `json:"containerPort,omitempty" yaml:"containerPort,omitempty"`

	// The name of the Cloud Run v2 Service.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
