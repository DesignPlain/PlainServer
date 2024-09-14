package types

type Sagemaker_SpaceSpaceSettingsJupyterServerAppSettings struct {
	// A list of Git repositories that SageMaker automatically displays to users for cloning in the JupyterServer application. See `code_repository` Block below.
	CodeRepositories []Sagemaker_SpaceSpaceSettingsJupyterServerAppSettingsCodeRepository `json:"codeRepositories,omitempty" yaml:"codeRepositories,omitempty"`

	// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. See `default_resource_spec` Block below.
	DefaultResourceSpec Sagemaker_SpaceSpaceSettingsJupyterServerAppSettingsDefaultResourceSpec `json:"defaultResourceSpec,omitempty" yaml:"defaultResourceSpec,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configurations.
	LifecycleConfigArns []string `json:"lifecycleConfigArns,omitempty" yaml:"lifecycleConfigArns,omitempty"`
}
