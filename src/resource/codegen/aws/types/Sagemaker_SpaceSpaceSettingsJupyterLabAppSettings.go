package types

type Sagemaker_SpaceSpaceSettingsJupyterLabAppSettings struct {
	// A list of Git repositories that SageMaker automatically displays to users for cloning in the JupyterServer application. See `code_repository` Block below.
	CodeRepositories []Sagemaker_SpaceSpaceSettingsJupyterLabAppSettingsCodeRepository `json:"codeRepositories,omitempty" yaml:"codeRepositories,omitempty"`

	// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. See `default_resource_spec` Block below.
	DefaultResourceSpec Sagemaker_SpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpec `json:"defaultResourceSpec,omitempty" yaml:"defaultResourceSpec,omitempty"`
}
