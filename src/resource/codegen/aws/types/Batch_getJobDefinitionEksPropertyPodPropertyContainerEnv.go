package types

type Batch_getJobDefinitionEksPropertyPodPropertyContainerEnv struct {
	// The name of the job definition to register. It can be up to 128 letters long. It can contain uppercase and lowercase letters, numbers, hyphens (-), and underscores (_).
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The quantity of the specified resource to reserve for the container.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
