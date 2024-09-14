package types

type Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerMountPoint struct {
	// The name of the volume to mount.
	SourceVolume string `json:"sourceVolume,omitempty" yaml:"sourceVolume,omitempty"`

	// The absolute file path in the container where the tmpfs volume is mounted.
	ContainerPath string `json:"containerPath,omitempty" yaml:"containerPath,omitempty"`

	// If this value is true, the container has read-only access to the volume.
	ReadOnly bool `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
}
