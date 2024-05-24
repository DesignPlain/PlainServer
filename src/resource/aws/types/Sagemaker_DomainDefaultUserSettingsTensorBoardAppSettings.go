package types

type Sagemaker_DomainDefaultUserSettingsTensorBoardAppSettings struct {
	// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see `default_resource_spec` Block below.
	DefaultResourceSpec Sagemaker_DomainDefaultUserSettingsTensorBoardAppSettingsDefaultResourceSpec `json:"defaultResourceSpec,omitempty" yaml:"defaultResourceSpec,omitempty"`
}
