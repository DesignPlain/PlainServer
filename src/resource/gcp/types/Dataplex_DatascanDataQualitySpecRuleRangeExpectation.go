package types

type Dataplex_DatascanDataQualitySpecRuleRangeExpectation struct {
	// The maximum column value allowed for a row to pass this validation. At least one of minValue and maxValue need to be provided.
	MaxValue string `json:"maxValue,omitempty" yaml:"maxValue,omitempty"`

	// The minimum column value allowed for a row to pass this validation. At least one of minValue and maxValue need to be provided.
	MinValue string `json:"minValue,omitempty" yaml:"minValue,omitempty"`

	/*
	   Whether each value needs to be strictly lesser than ('<') the maximum, or if equality is allowed.
	   Only relevant if a maxValue has been defined. Default = false.
	*/
	StrictMaxEnabled bool `json:"strictMaxEnabled,omitempty" yaml:"strictMaxEnabled,omitempty"`

	/*
	   Whether each value needs to be strictly greater than ('>') the minimum, or if equality is allowed.
	   Only relevant if a minValue has been defined. Default = false.
	*/
	StrictMinEnabled bool `json:"strictMinEnabled,omitempty" yaml:"strictMinEnabled,omitempty"`
}
