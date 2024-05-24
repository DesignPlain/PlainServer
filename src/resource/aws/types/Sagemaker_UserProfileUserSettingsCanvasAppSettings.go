package types

type Sagemaker_UserProfileUserSettingsCanvasAppSettings struct {
	// The model deployment settings for the SageMaker Canvas application. See Direct Deploy Settings below.
	DirectDeploySettings Sagemaker_UserProfileUserSettingsCanvasAppSettingsDirectDeploySettings `json:"directDeploySettings,omitempty" yaml:"directDeploySettings,omitempty"`

	// The settings for connecting to an external data source with OAuth. See Identity Provider OAuth Settings below.
	IdentityProviderOauthSettings []Sagemaker_UserProfileUserSettingsCanvasAppSettingsIdentityProviderOauthSetting `json:"identityProviderOauthSettings,omitempty" yaml:"identityProviderOauthSettings,omitempty"`

	// The settings for document querying. See Kendra Settings below.
	KendraSettings Sagemaker_UserProfileUserSettingsCanvasAppSettingsKendraSettings `json:"kendraSettings,omitempty" yaml:"kendraSettings,omitempty"`

	// The model registry settings for the SageMaker Canvas application. See Model Register Settings below.
	ModelRegisterSettings Sagemaker_UserProfileUserSettingsCanvasAppSettingsModelRegisterSettings `json:"modelRegisterSettings,omitempty" yaml:"modelRegisterSettings,omitempty"`

	// Time series forecast settings for the Canvas app. See Time Series Forecasting Settings below.
	TimeSeriesForecastingSettings Sagemaker_UserProfileUserSettingsCanvasAppSettingsTimeSeriesForecastingSettings `json:"timeSeriesForecastingSettings,omitempty" yaml:"timeSeriesForecastingSettings,omitempty"`

	// The workspace settings for the SageMaker Canvas application. See Workspace Settings below.
	WorkspaceSettings Sagemaker_UserProfileUserSettingsCanvasAppSettingsWorkspaceSettings `json:"workspaceSettings,omitempty" yaml:"workspaceSettings,omitempty"`
}
