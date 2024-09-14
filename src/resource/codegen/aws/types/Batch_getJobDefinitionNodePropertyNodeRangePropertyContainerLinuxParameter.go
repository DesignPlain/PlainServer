package types

type Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerLinuxParameter struct {
	// Any of the host devices to expose to the container.
	Devices []Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerLinuxParameterDevice `json:"devices,omitempty" yaml:"devices,omitempty"`

	// If true, run an init process inside the container that forwards signals and reaps processes.
	InitProcessEnabled bool `json:"initProcessEnabled,omitempty" yaml:"initProcessEnabled,omitempty"`

	// The total amount of swap memory (in MiB) a container can use.
	MaxSwap int `json:"maxSwap,omitempty" yaml:"maxSwap,omitempty"`

	// The value for the size (in MiB) of the `/dev/shm` volume.
	SharedMemorySize int `json:"sharedMemorySize,omitempty" yaml:"sharedMemorySize,omitempty"`

	// You can use this parameter to tune a container's memory swappiness behavior.
	Swappiness int `json:"swappiness,omitempty" yaml:"swappiness,omitempty"`

	// The container path, mount options, and size (in MiB) of the tmpfs mount.
	Tmpfs []Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerLinuxParameterTmpf `json:"tmpfs,omitempty" yaml:"tmpfs,omitempty"`
}
