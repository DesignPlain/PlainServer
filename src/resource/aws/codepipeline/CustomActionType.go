package codepipeline

import types "DesignSphere_Server/src/resource/aws/types"

type CustomActionType struct {
	// The category of the custom action. Valid values: `Source`, `Build`, `Deploy`, `Test`, `Invoke`, `Approval`
	Category string `json:"category,omitempty" yaml:"category,omitempty"`

	// The configuration properties for the custom action. Max 10 items.
	ConfigurationProperties []types.Codepipeline_CustomActionTypeConfigurationProperty `json:"configurationProperties,omitempty" yaml:"configurationProperties,omitempty"`

	// The details of the input artifact for the action.
	InputArtifactDetails types.Codepipeline_CustomActionTypeInputArtifactDetails `json:"inputArtifactDetails,omitempty" yaml:"inputArtifactDetails,omitempty"`

	// The details of the output artifact of the action.
	OutputArtifactDetails types.Codepipeline_CustomActionTypeOutputArtifactDetails `json:"outputArtifactDetails,omitempty" yaml:"outputArtifactDetails,omitempty"`

	// The provider of the service used in the custom action
	ProviderName string `json:"providerName,omitempty" yaml:"providerName,omitempty"`

	// The settings for an action type.
	Settings types.Codepipeline_CustomActionTypeSettings `json:"settings,omitempty" yaml:"settings,omitempty"`

	// Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The version identifier of the custom action.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
