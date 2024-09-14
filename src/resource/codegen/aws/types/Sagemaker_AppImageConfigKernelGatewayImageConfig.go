package types

type Sagemaker_AppImageConfigKernelGatewayImageConfig struct {
	// The URL where the Git repository is located. See File System Config details below.
	FileSystemConfig Sagemaker_AppImageConfigKernelGatewayImageConfigFileSystemConfig `json:"fileSystemConfig,omitempty" yaml:"fileSystemConfig,omitempty"`

	// The default branch for the Git repository. See Kernel Spec details below.
	KernelSpec Sagemaker_AppImageConfigKernelGatewayImageConfigKernelSpec `json:"kernelSpec,omitempty" yaml:"kernelSpec,omitempty"`
}
