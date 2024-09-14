package types

type Sagemaker_AppImageConfigCodeEditorAppImageConfig struct {
	// The configuration used to run the application image container. See Container Config details below.
	ContainerConfig Sagemaker_AppImageConfigCodeEditorAppImageConfigContainerConfig `json:"containerConfig,omitempty" yaml:"containerConfig,omitempty"`

	// The URL where the Git repository is located. See File System Config details below.
	FileSystemConfig Sagemaker_AppImageConfigCodeEditorAppImageConfigFileSystemConfig `json:"fileSystemConfig,omitempty" yaml:"fileSystemConfig,omitempty"`
}
