package types

type Quicksight_AnalysisParametersIntegerParameter struct {
	/*
	   Display name for the analysis.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Values []int `json:"values,omitempty" yaml:"values,omitempty"`
}
