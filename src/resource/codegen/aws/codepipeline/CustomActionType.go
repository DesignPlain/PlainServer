package codepipeline

import types "libds/aws/types"

type CustomActionType struct {
	//
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// The category of the custom action. Valid values: `Source`, `Build`, `Deploy`, `Test`, `Invoke`, `Approval`
	Category string `json:"category,omitempty" yaml:"category,omitempty"`

	// The configuration properties for the custom action. Max 10 items.
	ConfigurationProperties []types.Codepipeline_CustomActionTypeConfigurationProperty `json:"configurationProperties,omitempty" yaml:"configurationProperties,omitempty"`

	//
	InputArtifactDetails types.Codepipeline_CustomActionTypeInputArtifactDetails `json:"inputArtifactDetails,omitempty" yaml:"inputArtifactDetails,omitempty"`

	//
	OutputArtifactDetails types.Codepipeline_CustomActionTypeOutputArtifactDetails `json:"outputArtifactDetails,omitempty" yaml:"outputArtifactDetails,omitempty"`

	//
	ProviderName string `json:"providerName,omitempty" yaml:"providerName,omitempty"`

	//
	Settings types.Codepipeline_CustomActionTypeSettings `json:"settings,omitempty" yaml:"settings,omitempty"`
}
