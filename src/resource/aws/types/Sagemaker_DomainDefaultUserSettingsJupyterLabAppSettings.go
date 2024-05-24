package types

type Sagemaker_DomainDefaultUserSettingsJupyterLabAppSettings struct {
	// A list of Git repositories that SageMaker automatically displays to users for cloning in the JupyterServer application. see `code_repository` Block below.
	CodeRepositories []Sagemaker_DomainDefaultUserSettingsJupyterLabAppSettingsCodeRepository `json:"codeRepositories,omitempty" yaml:"codeRepositories,omitempty"`

	// A list of custom SageMaker images that are configured to run as a JupyterLab app. see `custom_image` Block below.
	CustomImages []Sagemaker_DomainDefaultUserSettingsJupyterLabAppSettingsCustomImage `json:"customImages,omitempty" yaml:"customImages,omitempty"`

	// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see `default_resource_spec` Block below.
	DefaultResourceSpec Sagemaker_DomainDefaultUserSettingsJupyterLabAppSettingsDefaultResourceSpec `json:"defaultResourceSpec,omitempty" yaml:"defaultResourceSpec,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configurations.
	LifecycleConfigArns []string `json:"lifecycleConfigArns,omitempty" yaml:"lifecycleConfigArns,omitempty"`
}
