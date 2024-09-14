package types

type Batch_JobDefinitionEksPropertiesPodPropertiesContainersVolumeMount struct {
	//
	MountPath string `json:"mountPath,omitempty" yaml:"mountPath,omitempty"`

	// Name of the job definition.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	ReadOnly bool `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
}
