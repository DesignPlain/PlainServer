package types

type Dataplex_DatascanDataQualitySpecRuleStatisticRangeExpectation struct {
	/*
	   The minimum column statistic value allowed for a row to pass this validation.
	   At least one of minValue and maxValue need to be provided.
	*/
	MinValue string `json:"minValue,omitempty" yaml:"minValue,omitempty"`

	/*
	   column statistics.
	   Possible values are: `STATISTIC_UNDEFINED`, `MEAN`, `MIN`, `MAX`.
	*/
	Statistic string `json:"statistic,omitempty" yaml:"statistic,omitempty"`

	/*
	   Whether column statistic needs to be strictly lesser than ('<') the maximum, or if equality is allowed.
	   Only relevant if a maxValue has been defined. Default = false.
	*/
	StrictMaxEnabled bool `json:"strictMaxEnabled,omitempty" yaml:"strictMaxEnabled,omitempty"`

	/*
	   Whether column statistic needs to be strictly greater than ('>') the minimum, or if equality is allowed.
	   Only relevant if a minValue has been defined. Default = false.
	*/
	StrictMinEnabled bool `json:"strictMinEnabled,omitempty" yaml:"strictMinEnabled,omitempty"`

	/*
	   The maximum column statistic value allowed for a row to pass this validation.
	   At least one of minValue and maxValue need to be provided.
	*/
	MaxValue string `json:"maxValue,omitempty" yaml:"maxValue,omitempty"`
}
