package sagemaker

import types "DesignSphere_Server/src/resource/aws/types"

type HumanTaskUI struct {
	// The Liquid template for the worker user interface. See UI Template below.
	UiTemplate types.Sagemaker_HumanTaskUIUiTemplate `json:"uiTemplate,omitempty" yaml:"uiTemplate,omitempty"`

	// The name of the Human Task UI.
	HumanTaskUiName string `json:"humanTaskUiName,omitempty" yaml:"humanTaskUiName,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
