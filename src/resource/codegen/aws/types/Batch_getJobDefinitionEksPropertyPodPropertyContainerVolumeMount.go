package types

type Batch_getJobDefinitionEksPropertyPodPropertyContainerVolumeMount struct {
	// The path on the container where the volume is mounted.
	MountPath string `json:"mountPath,omitempty" yaml:"mountPath,omitempty"`

	// The name of the job definition to register. It can be up to 128 letters long. It can contain uppercase and lowercase letters, numbers, hyphens (-), and underscores (_).
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// If this value is true, the container has read-only access to the volume.
	ReadOnly bool `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
}
