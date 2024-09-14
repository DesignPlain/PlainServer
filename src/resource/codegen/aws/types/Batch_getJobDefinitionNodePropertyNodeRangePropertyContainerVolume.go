package types

type Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerVolume struct {
	// The contents of the host parameter determine whether your data volume persists on the host container instance and where it's stored.
	Hosts []Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerVolumeHost `json:"hosts,omitempty" yaml:"hosts,omitempty"`

	// The name of the job definition to register. It can be up to 128 letters long. It can contain uppercase and lowercase letters, numbers, hyphens (-), and underscores (_).
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// This parameter is specified when you're using an Amazon Elastic File System file system for job storage.
	EfsVolumeConfigurations []Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerVolumeEfsVolumeConfiguration `json:"efsVolumeConfigurations,omitempty" yaml:"efsVolumeConfigurations,omitempty"`
}
