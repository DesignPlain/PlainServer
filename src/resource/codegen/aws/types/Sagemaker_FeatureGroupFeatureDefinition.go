package types

type Sagemaker_FeatureGroupFeatureDefinition struct {
	// The name of a feature. `feature_name` cannot be any of the following: `is_deleted`, `write_time`, `api_invocation_time`.
	FeatureName string `json:"featureName,omitempty" yaml:"featureName,omitempty"`

	// The value type of a feature. Valid values are `Integral`, `Fractional`, or `String`.
	FeatureType string `json:"featureType,omitempty" yaml:"featureType,omitempty"`
}
