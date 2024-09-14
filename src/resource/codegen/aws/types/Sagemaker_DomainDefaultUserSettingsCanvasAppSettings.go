package types

type Sagemaker_DomainDefaultUserSettingsCanvasAppSettings struct {
	//
	GenerativeAiSettings Sagemaker_DomainDefaultUserSettingsCanvasAppSettingsGenerativeAiSettings `json:"generativeAiSettings,omitempty" yaml:"generativeAiSettings,omitempty"`

	// The settings for connecting to an external data source with OAuth. See `identity_provider_oauth_settings` Block below.
	IdentityProviderOauthSettings []Sagemaker_DomainDefaultUserSettingsCanvasAppSettingsIdentityProviderOauthSetting `json:"identityProviderOauthSettings,omitempty" yaml:"identityProviderOauthSettings,omitempty"`

	// The settings for document querying. See `kendra_settings` Block below.
	KendraSettings Sagemaker_DomainDefaultUserSettingsCanvasAppSettingsKendraSettings `json:"kendraSettings,omitempty" yaml:"kendraSettings,omitempty"`

	// The model registry settings for the SageMaker Canvas application. See `model_register_settings` Block below.
	ModelRegisterSettings Sagemaker_DomainDefaultUserSettingsCanvasAppSettingsModelRegisterSettings `json:"modelRegisterSettings,omitempty" yaml:"modelRegisterSettings,omitempty"`

	// Time series forecast settings for the Canvas app. See `time_series_forecasting_settings` Block below.
	TimeSeriesForecastingSettings Sagemaker_DomainDefaultUserSettingsCanvasAppSettingsTimeSeriesForecastingSettings `json:"timeSeriesForecastingSettings,omitempty" yaml:"timeSeriesForecastingSettings,omitempty"`

	// The workspace settings for the SageMaker Canvas application. See `workspace_settings` Block below.
	WorkspaceSettings Sagemaker_DomainDefaultUserSettingsCanvasAppSettingsWorkspaceSettings `json:"workspaceSettings,omitempty" yaml:"workspaceSettings,omitempty"`

	// The model deployment settings for the SageMaker Canvas application. See `direct_deploy_settings` Block below.
	DirectDeploySettings Sagemaker_DomainDefaultUserSettingsCanvasAppSettingsDirectDeploySettings `json:"directDeploySettings,omitempty" yaml:"directDeploySettings,omitempty"`
}
