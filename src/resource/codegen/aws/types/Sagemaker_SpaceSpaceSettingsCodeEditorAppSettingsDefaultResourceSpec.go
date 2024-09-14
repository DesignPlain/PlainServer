package types

type Sagemaker_SpaceSpaceSettingsCodeEditorAppSettingsDefaultResourceSpec struct {
	// The instance type.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configuration attached to the Resource.
	LifecycleConfigArn string `json:"lifecycleConfigArn,omitempty" yaml:"lifecycleConfigArn,omitempty"`

	// The Amazon Resource Name (ARN) of the SageMaker image created on the instance.
	SagemakerImageArn string `json:"sagemakerImageArn,omitempty" yaml:"sagemakerImageArn,omitempty"`

	// The SageMaker Image Version Alias.
	SagemakerImageVersionAlias string `json:"sagemakerImageVersionAlias,omitempty" yaml:"sagemakerImageVersionAlias,omitempty"`

	// The ARN of the image version created on the instance.
	SagemakerImageVersionArn string `json:"sagemakerImageVersionArn,omitempty" yaml:"sagemakerImageVersionArn,omitempty"`
}
