package types

type Batch_JobDefinitionEksPropertiesPodPropertiesContainersEnv struct {
	// Specifies the name of the job definition.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The value of the environment variable.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
