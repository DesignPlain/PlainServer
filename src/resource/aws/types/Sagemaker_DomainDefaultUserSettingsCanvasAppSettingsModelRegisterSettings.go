package types

type Sagemaker_DomainDefaultUserSettingsCanvasAppSettingsModelRegisterSettings struct {
	// The Amazon Resource Name (ARN) of the SageMaker model registry account. Required only to register model versions created by a different SageMaker Canvas AWS account than the AWS account in which SageMaker model registry is set up.
	CrossAccountModelRegisterRoleArn string `json:"crossAccountModelRegisterRoleArn,omitempty" yaml:"crossAccountModelRegisterRoleArn,omitempty"`

	// Describes whether the integration to the model registry is enabled or disabled in the Canvas application. Valid values are `ENABLED` and `DISABLED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
