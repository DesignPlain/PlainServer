package types

type Sagemaker_SpaceSpaceSettings struct {
	// The Code Editor application settings. See `code_editor_app_settings` Block below.
	CodeEditorAppSettings Sagemaker_SpaceSpaceSettingsCodeEditorAppSettings `json:"codeEditorAppSettings,omitempty" yaml:"codeEditorAppSettings,omitempty"`

	// A file system, created by you, that you assign to a space for an Amazon SageMaker Domain. See `custom_file_system` Block below.
	CustomFileSystems []Sagemaker_SpaceSpaceSettingsCustomFileSystem `json:"customFileSystems,omitempty" yaml:"customFileSystems,omitempty"`

	// The settings for the JupyterLab application. See `jupyter_lab_app_settings` Block below.
	JupyterLabAppSettings Sagemaker_SpaceSpaceSettingsJupyterLabAppSettings `json:"jupyterLabAppSettings,omitempty" yaml:"jupyterLabAppSettings,omitempty"`

	// The Jupyter server's app settings. See `jupyter_server_app_settings` Block below.
	JupyterServerAppSettings Sagemaker_SpaceSpaceSettingsJupyterServerAppSettings `json:"jupyterServerAppSettings,omitempty" yaml:"jupyterServerAppSettings,omitempty"`

	// The kernel gateway app settings. See `kernel_gateway_app_settings` Block below.
	KernelGatewayAppSettings Sagemaker_SpaceSpaceSettingsKernelGatewayAppSettings `json:"kernelGatewayAppSettings,omitempty" yaml:"kernelGatewayAppSettings,omitempty"`

	// The storage settings. See `space_storage_settings` Block below.
	SpaceStorageSettings Sagemaker_SpaceSpaceSettingsSpaceStorageSettings `json:"spaceStorageSettings,omitempty" yaml:"spaceStorageSettings,omitempty"`

	// The type of app created within the space.
	AppType string `json:"appType,omitempty" yaml:"appType,omitempty"`
}
