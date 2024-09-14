package types

type Batch_JobDefinitionEksPropertiesPodPropertiesContainersSecurityContext struct {
	//
	Privileged bool `json:"privileged,omitempty" yaml:"privileged,omitempty"`

	//
	ReadOnlyRootFileSystem bool `json:"readOnlyRootFileSystem,omitempty" yaml:"readOnlyRootFileSystem,omitempty"`

	//
	RunAsGroup int `json:"runAsGroup,omitempty" yaml:"runAsGroup,omitempty"`

	//
	RunAsNonRoot bool `json:"runAsNonRoot,omitempty" yaml:"runAsNonRoot,omitempty"`

	//
	RunAsUser int `json:"runAsUser,omitempty" yaml:"runAsUser,omitempty"`
}
