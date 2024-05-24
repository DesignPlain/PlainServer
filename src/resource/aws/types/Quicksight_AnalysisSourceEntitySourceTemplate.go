package types

type Quicksight_AnalysisSourceEntitySourceTemplate struct {
	// The Amazon Resource Name (ARN) of the resource.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// List of dataset references. See data_set_references.
	DataSetReferences []Quicksight_AnalysisSourceEntitySourceTemplateDataSetReference `json:"dataSetReferences,omitempty" yaml:"dataSetReferences,omitempty"`
}
