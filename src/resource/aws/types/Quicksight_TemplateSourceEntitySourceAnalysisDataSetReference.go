package types

type Quicksight_TemplateSourceEntitySourceAnalysisDataSetReference struct {
	// Dataset placeholder.
	DataSetPlaceholder string `json:"dataSetPlaceholder,omitempty" yaml:"dataSetPlaceholder,omitempty"`

	// Dataset Amazon Resource Name (ARN).
	DataSetArn string `json:"dataSetArn,omitempty" yaml:"dataSetArn,omitempty"`
}
