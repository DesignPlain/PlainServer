package types

type Securityhub_InsightFiltersResourceDetailsOther struct {
	//
	Comparison string `json:"comparison,omitempty" yaml:"comparison,omitempty"`

	// The key of the map filter. For example, for `ResourceTags`, `Key` identifies the name of the tag. For `UserDefinedFields`, `Key` is the name of the field.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	//
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
