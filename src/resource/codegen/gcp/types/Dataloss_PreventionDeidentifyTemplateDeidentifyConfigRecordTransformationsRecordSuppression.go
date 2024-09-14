package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppression struct {
	/*
	   A condition that when it evaluates to true will result in the record being evaluated to be suppressed from the transformed content.
	   Structure is documented below.
	*/
	Condition Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionCondition `json:"condition,omitempty" yaml:"condition,omitempty"`
}
