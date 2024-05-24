package types

type Sagemaker_AppImageConfigJupyterLabImageConfigContainerConfig struct {
	// The environment variables to set in the container.
	ContainerEnvironmentVariables map[string]string `json:"containerEnvironmentVariables,omitempty" yaml:"containerEnvironmentVariables,omitempty"`

	// The arguments for the container when you're running the application.
	ContainerArguments []string `json:"containerArguments,omitempty" yaml:"containerArguments,omitempty"`

	// The entrypoint used to run the application in the container.
	ContainerEntrypoints []string `json:"containerEntrypoints,omitempty" yaml:"containerEntrypoints,omitempty"`
}
