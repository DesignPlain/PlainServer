package types

type Sagemaker_AppImageConfigJupyterLabImageConfig struct {
	// The configuration used to run the application image container. See Container Config details below.
	ContainerConfig Sagemaker_AppImageConfigJupyterLabImageConfigContainerConfig `json:"containerConfig,omitempty" yaml:"containerConfig,omitempty"`
}
