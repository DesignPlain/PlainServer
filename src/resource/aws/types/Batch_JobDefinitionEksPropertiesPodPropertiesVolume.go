package types

type Batch_JobDefinitionEksPropertiesPodPropertiesVolume struct {
	//
	EmptyDir Batch_JobDefinitionEksPropertiesPodPropertiesVolumeEmptyDir `json:"emptyDir,omitempty" yaml:"emptyDir,omitempty"`

	//
	HostPath Batch_JobDefinitionEksPropertiesPodPropertiesVolumeHostPath `json:"hostPath,omitempty" yaml:"hostPath,omitempty"`

	// Specifies the name of the job definition.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Secret Batch_JobDefinitionEksPropertiesPodPropertiesVolumeSecret `json:"secret,omitempty" yaml:"secret,omitempty"`
}
