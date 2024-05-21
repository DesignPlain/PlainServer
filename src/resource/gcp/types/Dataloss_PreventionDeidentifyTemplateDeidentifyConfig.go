package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfig struct {
	/*
	   Treat the dataset as an image and redact.
	   Structure is documented below.
	*/
	ImageTransformations Dataloss_PreventionDeidentifyTemplateDeidentifyConfigImageTransformations `json:"imageTransformations,omitempty" yaml:"imageTransformations,omitempty"`

	/*
	   Treat the dataset as free-form text and apply the same free text transformation everywhere
	   Structure is documented below.
	*/
	InfoTypeTransformations Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformations `json:"infoTypeTransformations,omitempty" yaml:"infoTypeTransformations,omitempty"`

	/*
	   Treat the dataset as structured. Transformations can be applied to specific locations within structured datasets, such as transforming a column within a table.
	   Structure is documented below.
	*/
	RecordTransformations Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformations `json:"recordTransformations,omitempty" yaml:"recordTransformations,omitempty"`
}
