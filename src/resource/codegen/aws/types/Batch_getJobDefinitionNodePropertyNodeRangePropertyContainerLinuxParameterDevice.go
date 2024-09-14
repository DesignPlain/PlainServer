package types

type Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerLinuxParameterDevice struct {
	// The absolute file path in the container where the tmpfs volume is mounted.
	ContainerPath string `json:"containerPath,omitempty" yaml:"containerPath,omitempty"`

	// The path for the device on the host container instance.
	HostPath string `json:"hostPath,omitempty" yaml:"hostPath,omitempty"`

	// The explicit permissions to provide to the container for the device.
	Permissions []string `json:"permissions,omitempty" yaml:"permissions,omitempty"`
}
