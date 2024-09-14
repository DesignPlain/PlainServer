package evidently

import types "libds/aws/types"

type Feature struct {
	// Tags to apply to the feature. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// One or more blocks that contain the configuration of the feature's different variations. Detailed below
	Variations []types.Evidently_FeatureVariation `json:"variations,omitempty" yaml:"variations,omitempty"`

	// The name of the variation to use as the default variation. The default variation is served to users who are not allocated to any ongoing launches or experiments of this feature. This variation must also be listed in the `variations` structure. If you omit `default_variation`, the first variation listed in the `variations` structure is used as the default variation.
	DefaultVariation string `json:"defaultVariation,omitempty" yaml:"defaultVariation,omitempty"`

	// Specifies the description of the feature.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Specify users that should always be served a specific variation of a feature. Each user is specified by a key-value pair . For each key, specify a user by entering their user ID, account ID, or some other identifier. For the value, specify the name of the variation that they are to be served.
	EntityOverrides map[string]string `json:"entityOverrides,omitempty" yaml:"entityOverrides,omitempty"`

	// Specify `ALL_RULES` to activate the traffic allocation specified by any ongoing launches or experiments. Specify `DEFAULT_VARIATION` to serve the default variation to all users instead.
	EvaluationStrategy string `json:"evaluationStrategy,omitempty" yaml:"evaluationStrategy,omitempty"`

	// The name for the new feature. Minimum length of `1`. Maximum length of `127`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The name or ARN of the project that is to contain the new feature.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
