package types

type Quicksight_TemplateSourceEntity struct {
	// The source template, if it is based on an template.. Only one of `source_analysis` or `source_template` should be configured. See source_template.
	SourceTemplate Quicksight_TemplateSourceEntitySourceTemplate `json:"sourceTemplate,omitempty" yaml:"sourceTemplate,omitempty"`

	// The source analysis, if it is based on an analysis.. Only one of `source_analysis` or `source_template` should be configured. See source_analysis.
	SourceAnalysis Quicksight_TemplateSourceEntitySourceAnalysis `json:"sourceAnalysis,omitempty" yaml:"sourceAnalysis,omitempty"`
}
