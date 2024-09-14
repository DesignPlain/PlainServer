package types

type Sagemaker_AppImageConfigJupyterLabImageConfig struct {
	// The URL where the Git repository is located. See File System Config details below.
	FileSystemConfig Sagemaker_AppImageConfigJupyterLabImageConfigFileSystemConfig `json:"fileSystemConfig,omitempty" yaml:"fileSystemConfig,omitempty"`

	// The configuration used to run the application image container. See Container Config details below.
	ContainerConfig Sagemaker_AppImageConfigJupyterLabImageConfigContainerConfig `json:"containerConfig,omitempty" yaml:"containerConfig,omitempty"`
}
