package types

type Batch_getJobDefinitionEksPropertyPodPropertyVolume struct {
	// The name of the job definition to register. It can be up to 128 letters long. It can contain uppercase and lowercase letters, numbers, hyphens (-), and underscores (_).
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Specifies the configuration of a Kubernetes secret volume.
	Secrets []Batch_getJobDefinitionEksPropertyPodPropertyVolumeSecret `json:"secrets,omitempty" yaml:"secrets,omitempty"`

	// Specifies the configuration of a Kubernetes emptyDir volume.
	EmptyDirs []Batch_getJobDefinitionEksPropertyPodPropertyVolumeEmptyDir `json:"emptyDirs,omitempty" yaml:"emptyDirs,omitempty"`

	// The path for the device on the host container instance.
	HostPaths []Batch_getJobDefinitionEksPropertyPodPropertyVolumeHostPath `json:"hostPaths,omitempty" yaml:"hostPaths,omitempty"`
}
