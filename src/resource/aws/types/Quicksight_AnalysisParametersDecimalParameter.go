package types

type Quicksight_AnalysisParametersDecimalParameter struct {
	/*
	   Display name for the analysis.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Values []float64 `json:"values,omitempty" yaml:"values,omitempty"`
}
