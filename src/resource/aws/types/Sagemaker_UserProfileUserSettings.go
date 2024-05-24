package types

type Sagemaker_UserProfileUserSettings struct {
	// The TensorBoard app settings. See TensorBoard App Settings below.
	TensorBoardAppSettings Sagemaker_UserProfileUserSettingsTensorBoardAppSettings `json:"tensorBoardAppSettings,omitempty" yaml:"tensorBoardAppSettings,omitempty"`

	// A collection of settings that configure user interaction with the RStudioServerPro app. See RStudioServerProAppSettings below.
	RStudioServerProAppSettings Sagemaker_UserProfileUserSettingsRStudioServerProAppSettings `json:"rStudioServerProAppSettings,omitempty" yaml:"rStudioServerProAppSettings,omitempty"`

	// A list of security group IDs that will be attached to the user.
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	// The execution role ARN for the user.
	ExecutionRole string `json:"executionRole,omitempty" yaml:"executionRole,omitempty"`

	// The RSession app settings. See RSession App Settings below.
	RSessionAppSettings Sagemaker_UserProfileUserSettingsRSessionAppSettings `json:"rSessionAppSettings,omitempty" yaml:"rSessionAppSettings,omitempty"`

	// The sharing settings. See Sharing Settings below.
	SharingSettings Sagemaker_UserProfileUserSettingsSharingSettings `json:"sharingSettings,omitempty" yaml:"sharingSettings,omitempty"`

	// The storage settings for a private space. See Space Storage Settings below.
	SpaceStorageSettings Sagemaker_UserProfileUserSettingsSpaceStorageSettings `json:"spaceStorageSettings,omitempty" yaml:"spaceStorageSettings,omitempty"`

	// Details about the POSIX identity that is used for file system operations. See Custom Posix User Config below.
	CustomPosixUserConfig Sagemaker_UserProfileUserSettingsCustomPosixUserConfig `json:"customPosixUserConfig,omitempty" yaml:"customPosixUserConfig,omitempty"`

	// The default experience that the user is directed to when accessing the domain. The supported values are: `studio::`: Indicates that Studio is the default experience. This value can only be passed if StudioWebPortal is set to ENABLED. `app:JupyterServer:`: Indicates that Studio Classic is the default experience.
	DefaultLandingUri string `json:"defaultLandingUri,omitempty" yaml:"defaultLandingUri,omitempty"`

	// The settings for assigning a custom file system to a user profile. Permitted users can access this file system in Amazon SageMaker Studio. See Custom File System Config below.
	CustomFileSystemConfigs []Sagemaker_UserProfileUserSettingsCustomFileSystemConfig `json:"customFileSystemConfigs,omitempty" yaml:"customFileSystemConfigs,omitempty"`

	// Whether the user can access Studio. If this value is set to `DISABLED`, the user cannot access Studio, even if that is the default experience for the domain. Valid values are `ENABLED` and `DISABLED`.
	StudioWebPortal string `json:"studioWebPortal,omitempty" yaml:"studioWebPortal,omitempty"`

	// The settings for the JupyterLab application. See Jupyter Lab App Settings below.
	JupyterLabAppSettings Sagemaker_UserProfileUserSettingsJupyterLabAppSettings `json:"jupyterLabAppSettings,omitempty" yaml:"jupyterLabAppSettings,omitempty"`

	// The Jupyter server's app settings. See Jupyter Server App Settings below.
	JupyterServerAppSettings Sagemaker_UserProfileUserSettingsJupyterServerAppSettings `json:"jupyterServerAppSettings,omitempty" yaml:"jupyterServerAppSettings,omitempty"`

	// The kernel gateway app settings. See Kernel Gateway App Settings below.
	KernelGatewayAppSettings Sagemaker_UserProfileUserSettingsKernelGatewayAppSettings `json:"kernelGatewayAppSettings,omitempty" yaml:"kernelGatewayAppSettings,omitempty"`

	// The Canvas app settings. See Canvas App Settings below.
	CanvasAppSettings Sagemaker_UserProfileUserSettingsCanvasAppSettings `json:"canvasAppSettings,omitempty" yaml:"canvasAppSettings,omitempty"`

	// The Code Editor application settings. See Code Editor App Settings below.
	CodeEditorAppSettings Sagemaker_UserProfileUserSettingsCodeEditorAppSettings `json:"codeEditorAppSettings,omitempty" yaml:"codeEditorAppSettings,omitempty"`
}
