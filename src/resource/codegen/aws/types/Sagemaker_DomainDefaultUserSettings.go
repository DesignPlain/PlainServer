package types

type Sagemaker_DomainDefaultUserSettings struct {
	// A list of security group IDs that will be attached to the user.
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	// The Canvas app settings. See `canvas_app_settings` Block below.
	CanvasAppSettings Sagemaker_DomainDefaultUserSettingsCanvasAppSettings `json:"canvasAppSettings,omitempty" yaml:"canvasAppSettings,omitempty"`

	// The settings for assigning a custom file system to a user profile. Permitted users can access this file system in Amazon SageMaker Studio. See `custom_file_system_config` Block below.
	CustomFileSystemConfigs []Sagemaker_DomainDefaultUserSettingsCustomFileSystemConfig `json:"customFileSystemConfigs,omitempty" yaml:"customFileSystemConfigs,omitempty"`

	// The settings for the JupyterLab application. See `jupyter_lab_app_settings` Block below.
	JupyterLabAppSettings Sagemaker_DomainDefaultUserSettingsJupyterLabAppSettings `json:"jupyterLabAppSettings,omitempty" yaml:"jupyterLabAppSettings,omitempty"`

	// A collection of settings that configure user interaction with the RStudioServerPro app. See `r_studio_server_pro_app_settings` Block below.
	RStudioServerProAppSettings Sagemaker_DomainDefaultUserSettingsRStudioServerProAppSettings `json:"rStudioServerProAppSettings,omitempty" yaml:"rStudioServerProAppSettings,omitempty"`

	// The storage settings for a private space. See `space_storage_settings` Block below.
	SpaceStorageSettings Sagemaker_DomainDefaultUserSettingsSpaceStorageSettings `json:"spaceStorageSettings,omitempty" yaml:"spaceStorageSettings,omitempty"`

	// Details about the POSIX identity that is used for file system operations. See `custom_posix_user_config` Block below.
	CustomPosixUserConfig Sagemaker_DomainDefaultUserSettingsCustomPosixUserConfig `json:"customPosixUserConfig,omitempty" yaml:"customPosixUserConfig,omitempty"`

	// The default experience that the user is directed to when accessing the domain. The supported values are: `studio::`: Indicates that Studio is the default experience. This value can only be passed if StudioWebPortal is set to ENABLED. `app:JupyterServer:`: Indicates that Studio Classic is the default experience.
	DefaultLandingUri string `json:"defaultLandingUri,omitempty" yaml:"defaultLandingUri,omitempty"`

	// The kernel gateway app settings. See `kernel_gateway_app_settings` Block below.
	KernelGatewayAppSettings Sagemaker_DomainDefaultUserSettingsKernelGatewayAppSettings `json:"kernelGatewayAppSettings,omitempty" yaml:"kernelGatewayAppSettings,omitempty"`

	// The RSession app settings. See `r_session_app_settings` Block below.
	RSessionAppSettings Sagemaker_DomainDefaultUserSettingsRSessionAppSettings `json:"rSessionAppSettings,omitempty" yaml:"rSessionAppSettings,omitempty"`

	// The sharing settings. See `sharing_settings` Block below.
	SharingSettings Sagemaker_DomainDefaultUserSettingsSharingSettings `json:"sharingSettings,omitempty" yaml:"sharingSettings,omitempty"`

	// Whether the user can access Studio. If this value is set to `DISABLED`, the user cannot access Studio, even if that is the default experience for the domain. Valid values are `ENABLED` and `DISABLED`.
	StudioWebPortal string `json:"studioWebPortal,omitempty" yaml:"studioWebPortal,omitempty"`

	// The Code Editor application settings. See `code_editor_app_settings` Block below.
	CodeEditorAppSettings Sagemaker_DomainDefaultUserSettingsCodeEditorAppSettings `json:"codeEditorAppSettings,omitempty" yaml:"codeEditorAppSettings,omitempty"`

	// The execution role ARN for the user.
	ExecutionRole string `json:"executionRole,omitempty" yaml:"executionRole,omitempty"`

	// The Jupyter server's app settings. See `jupyter_server_app_settings` Block below.
	JupyterServerAppSettings Sagemaker_DomainDefaultUserSettingsJupyterServerAppSettings `json:"jupyterServerAppSettings,omitempty" yaml:"jupyterServerAppSettings,omitempty"`

	// The TensorBoard app settings. See `tensor_board_app_settings` Block below.
	TensorBoardAppSettings Sagemaker_DomainDefaultUserSettingsTensorBoardAppSettings `json:"tensorBoardAppSettings,omitempty" yaml:"tensorBoardAppSettings,omitempty"`
}
