package types

type Sagemaker_DomainDefaultUserSettingsCodeEditorAppSettings struct {
	// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see `default_resource_spec` Block below.
	DefaultResourceSpec Sagemaker_DomainDefaultUserSettingsCodeEditorAppSettingsDefaultResourceSpec `json:"defaultResourceSpec,omitempty" yaml:"defaultResourceSpec,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configurations.
	LifecycleConfigArns []string `json:"lifecycleConfigArns,omitempty" yaml:"lifecycleConfigArns,omitempty"`

	// A list of custom SageMaker images that are configured to run as a CodeEditor app. see `custom_image` Block below.
	CustomImages []Sagemaker_DomainDefaultUserSettingsCodeEditorAppSettingsCustomImage `json:"customImages,omitempty" yaml:"customImages,omitempty"`
}
