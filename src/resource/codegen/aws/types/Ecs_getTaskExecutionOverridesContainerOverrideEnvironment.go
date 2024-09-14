package types

type Ecs_getTaskExecutionOverridesContainerOverrideEnvironment struct {
	// The value of the key-value pair. For environment variables, this is the value of the environment variable.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// The name of the key-value pair. For environment variables, this is the name of the environment variable.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
