package types

type Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerLinuxParameterTmpf struct {
	// The list of tmpfs volume mount options.
	MountOptions []string `json:"mountOptions,omitempty" yaml:"mountOptions,omitempty"`

	// The size (in MiB) of the tmpfs volume.
	Size int `json:"size,omitempty" yaml:"size,omitempty"`

	// The absolute file path in the container where the tmpfs volume is mounted.
	ContainerPath string `json:"containerPath,omitempty" yaml:"containerPath,omitempty"`
}
