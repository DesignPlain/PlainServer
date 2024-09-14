package types

type Sagemaker_DomainDefaultSpaceSettings struct {
	// The execution role for the space.
	ExecutionRole string `json:"executionRole,omitempty" yaml:"executionRole,omitempty"`

	// The Jupyter server's app settings. See `jupyter_server_app_settings` Block below.
	JupyterServerAppSettings Sagemaker_DomainDefaultSpaceSettingsJupyterServerAppSettings `json:"jupyterServerAppSettings,omitempty" yaml:"jupyterServerAppSettings,omitempty"`

	// The kernel gateway app settings. See `kernel_gateway_app_settings` Block below.
	KernelGatewayAppSettings Sagemaker_DomainDefaultSpaceSettingsKernelGatewayAppSettings `json:"kernelGatewayAppSettings,omitempty" yaml:"kernelGatewayAppSettings,omitempty"`

	// The security groups for the Amazon Virtual Private Cloud that the space uses for communication.
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`
}
