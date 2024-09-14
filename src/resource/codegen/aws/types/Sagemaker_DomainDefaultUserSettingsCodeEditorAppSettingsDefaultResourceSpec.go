package types

type Sagemaker_DomainDefaultUserSettingsCodeEditorAppSettingsDefaultResourceSpec struct {
	// The SageMaker Image Version Alias.
	SagemakerImageVersionAlias string `json:"sagemakerImageVersionAlias,omitempty" yaml:"sagemakerImageVersionAlias,omitempty"`

	// The ARN of the image version created on the instance.
	SagemakerImageVersionArn string `json:"sagemakerImageVersionArn,omitempty" yaml:"sagemakerImageVersionArn,omitempty"`

	// The instance type that the image version runs on.. For valid values see [SageMaker Instance Types](https://docs.aws.amazon.com/sagemaker/latest/dg/notebooks-available-instance-types.html).
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configuration attached to the Resource.
	LifecycleConfigArn string `json:"lifecycleConfigArn,omitempty" yaml:"lifecycleConfigArn,omitempty"`

	// The ARN of the SageMaker image that the image version belongs to.
	SagemakerImageArn string `json:"sagemakerImageArn,omitempty" yaml:"sagemakerImageArn,omitempty"`
}
