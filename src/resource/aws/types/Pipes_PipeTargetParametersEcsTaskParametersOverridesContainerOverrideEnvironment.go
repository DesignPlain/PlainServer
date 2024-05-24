package types

type Pipes_PipeTargetParametersEcsTaskParametersOverridesContainerOverrideEnvironment struct {
	// Name of the pipe. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Value of parameter to start execution of a SageMaker Model Building Pipeline. Maximum length of 1024.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
