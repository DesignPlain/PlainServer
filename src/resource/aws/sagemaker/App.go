package sagemaker

import types "DesignSphere_Server/src/resource/aws/types"

type App struct {
	// The type of app. Valid values are `JupyterServer`, `KernelGateway`, `RStudioServerPro`, `RSessionGateway` and `TensorBoard`.
	AppType string `json:"appType,omitempty" yaml:"appType,omitempty"`

	// The domain ID.
	DomainId string `json:"domainId,omitempty" yaml:"domainId,omitempty"`

	// The instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance.See Resource Spec below.
	ResourceSpec types.Sagemaker_AppResourceSpec `json:"resourceSpec,omitempty" yaml:"resourceSpec,omitempty"`

	// The name of the space. At least one of `user_profile_name` or `space_name` required.
	SpaceName string `json:"spaceName,omitempty" yaml:"spaceName,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The user profile name. At least one of `user_profile_name` or `space_name` required.
	UserProfileName string `json:"userProfileName,omitempty" yaml:"userProfileName,omitempty"`

	// The name of the app.
	AppName string `json:"appName,omitempty" yaml:"appName,omitempty"`
}
