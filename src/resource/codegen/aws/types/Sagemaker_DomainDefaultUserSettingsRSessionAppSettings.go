package types

type Sagemaker_DomainDefaultUserSettingsRSessionAppSettings struct {
	// A list of custom SageMaker images that are configured to run as a RSession app. see `custom_image` Block below.
	CustomImages []Sagemaker_DomainDefaultUserSettingsRSessionAppSettingsCustomImage `json:"customImages,omitempty" yaml:"customImages,omitempty"`

	// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see `default_resource_spec` Block above.
	DefaultResourceSpec Sagemaker_DomainDefaultUserSettingsRSessionAppSettingsDefaultResourceSpec `json:"defaultResourceSpec,omitempty" yaml:"defaultResourceSpec,omitempty"`
}
