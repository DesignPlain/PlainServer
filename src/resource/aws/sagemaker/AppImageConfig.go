package sagemaker

import types "DesignSphere_Server/src/resource/aws/types"

type AppImageConfig struct {
	// The name of the App Image Config.
	AppImageConfigName string `json:"appImageConfigName,omitempty" yaml:"appImageConfigName,omitempty"`

	//
	JupyterLabImageConfig types.Sagemaker_AppImageConfigJupyterLabImageConfig `json:"jupyterLabImageConfig,omitempty" yaml:"jupyterLabImageConfig,omitempty"`

	// The configuration for the file system and kernels in a SageMaker image running as a KernelGateway app. See Kernel Gateway Image Config details below.
	KernelGatewayImageConfig types.Sagemaker_AppImageConfigKernelGatewayImageConfig `json:"kernelGatewayImageConfig,omitempty" yaml:"kernelGatewayImageConfig,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
