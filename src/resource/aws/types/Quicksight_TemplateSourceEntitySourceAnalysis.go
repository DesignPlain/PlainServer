package types

type Quicksight_TemplateSourceEntitySourceAnalysis struct {
	// The Amazon Resource Name (ARN) of the resource.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// A list of dataset references used as placeholders in the template. See data_set_references.
	DataSetReferences []Quicksight_TemplateSourceEntitySourceAnalysisDataSetReference `json:"dataSetReferences,omitempty" yaml:"dataSetReferences,omitempty"`
}
