package types

type Evidently_FeatureVariation struct {
	// A block that specifies the value assigned to this variation. Detailed below
	Value Evidently_FeatureVariationValue `json:"value,omitempty" yaml:"value,omitempty"`

	// The name of the variation. Minimum length of `1`. Maximum length of `127`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
