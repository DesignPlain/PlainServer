package types

type Quicksight_AnalysisParametersStringParameter struct {
	/*
	   Display name for the analysis.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
