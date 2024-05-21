package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformations struct {
	/*
	   Transform the record by applying various field transformations.
	   Structure is documented below.
	*/
	FieldTransformations []Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformation `json:"fieldTransformations,omitempty" yaml:"fieldTransformations,omitempty"`

	/*
	   Configuration defining which records get suppressed entirely. Records that match any suppression rule are omitted from the output.
	   Structure is documented below.
	*/
	RecordSuppressions []Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppression `json:"recordSuppressions,omitempty" yaml:"recordSuppressions,omitempty"`
}
