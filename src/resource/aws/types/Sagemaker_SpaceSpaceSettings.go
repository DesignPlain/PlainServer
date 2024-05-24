package types

type Sagemaker_SpaceSpaceSettings struct {
	//
	SpaceStorageSettings Sagemaker_SpaceSpaceSettingsSpaceStorageSettings `json:"spaceStorageSettings,omitempty" yaml:"spaceStorageSettings,omitempty"`

	// The type of app created within the space.
	AppType string `json:"appType,omitempty" yaml:"appType,omitempty"`

	// The Code Editor application settings. See Code Editor App Settings below.
	CodeEditorAppSettings Sagemaker_SpaceSpaceSettingsCodeEditorAppSettings `json:"codeEditorAppSettings,omitempty" yaml:"codeEditorAppSettings,omitempty"`

	// A file system, created by you, that you assign to a space for an Amazon SageMaker Domain. See Custom File System below.
	CustomFileSystems []Sagemaker_SpaceSpaceSettingsCustomFileSystem `json:"customFileSystems,omitempty" yaml:"customFileSystems,omitempty"`

	// The settings for the JupyterLab application. See Jupyter Lab App Settings below.
	JupyterLabAppSettings Sagemaker_SpaceSpaceSettingsJupyterLabAppSettings `json:"jupyterLabAppSettings,omitempty" yaml:"jupyterLabAppSettings,omitempty"`

	// The Jupyter server's app settings. See Jupyter Server App Settings below.
	JupyterServerAppSettings Sagemaker_SpaceSpaceSettingsJupyterServerAppSettings `json:"jupyterServerAppSettings,omitempty" yaml:"jupyterServerAppSettings,omitempty"`

	// The kernel gateway app settings. See Kernel Gateway App Settings below.
	KernelGatewayAppSettings Sagemaker_SpaceSpaceSettingsKernelGatewayAppSettings `json:"kernelGatewayAppSettings,omitempty" yaml:"kernelGatewayAppSettings,omitempty"`
}
