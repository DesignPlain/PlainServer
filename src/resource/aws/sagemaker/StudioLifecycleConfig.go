package sagemaker

type StudioLifecycleConfig struct {
	// The App type that the Lifecycle Configuration is attached to. Valid values are `JupyterServer`, `JupyterLab`, `CodeEditor` and `KernelGateway`.
	StudioLifecycleConfigAppType string `json:"studioLifecycleConfigAppType,omitempty" yaml:"studioLifecycleConfigAppType,omitempty"`

	// The content of your Studio Lifecycle Configuration script. This content must be base64 encoded.
	StudioLifecycleConfigContent string `json:"studioLifecycleConfigContent,omitempty" yaml:"studioLifecycleConfigContent,omitempty"`

	// The name of the Studio Lifecycle Configuration to create.
	StudioLifecycleConfigName string `json:"studioLifecycleConfigName,omitempty" yaml:"studioLifecycleConfigName,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
