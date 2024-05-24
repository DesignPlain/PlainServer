package types

type Evidently_FeatureVariationValue struct {
	// If this feature uses the long variation type, this field contains the long value of this variation. Minimum value of `-9007199254740991`. Maximum value of `9007199254740991`.
	LongValue string `json:"longValue,omitempty" yaml:"longValue,omitempty"`

	// If this feature uses the string variation type, this field contains the string value of this variation. Minimum length of `0`. Maximum length of `512`.
	StringValue string `json:"stringValue,omitempty" yaml:"stringValue,omitempty"`

	// If this feature uses the Boolean variation type, this field contains the Boolean value of this variation.
	BoolValue string `json:"boolValue,omitempty" yaml:"boolValue,omitempty"`

	// If this feature uses the double integer variation type, this field contains the double integer value of this variation.
	DoubleValue string `json:"doubleValue,omitempty" yaml:"doubleValue,omitempty"`
}
