package types

type Sagemaker_DomainDefaultSpaceSettingsJupyterServerAppSettings struct {
	// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see `default_resource_spec` Block below.
	DefaultResourceSpec Sagemaker_DomainDefaultSpaceSettingsJupyterServerAppSettingsDefaultResourceSpec `json:"defaultResourceSpec,omitempty" yaml:"defaultResourceSpec,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configurations.
	LifecycleConfigArns []string `json:"lifecycleConfigArns,omitempty" yaml:"lifecycleConfigArns,omitempty"`

	// A list of Git repositories that SageMaker automatically displays to users for cloning in the JupyterServer application. see `code_repository` Block below.
	CodeRepositories []Sagemaker_DomainDefaultSpaceSettingsJupyterServerAppSettingsCodeRepository `json:"codeRepositories,omitempty" yaml:"codeRepositories,omitempty"`
}
