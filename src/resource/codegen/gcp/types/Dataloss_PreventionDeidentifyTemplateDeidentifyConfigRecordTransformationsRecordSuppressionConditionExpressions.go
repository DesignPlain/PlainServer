package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionConditionExpressions struct {
	/*
	   Conditions to apply to the expression.
	   Structure is documented below.
	*/
	Conditions Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionConditionExpressionsConditions `json:"conditions,omitempty" yaml:"conditions,omitempty"`

	/*
	   The operator to apply to the result of conditions. Default and currently only supported value is AND.
	   Default value is `AND`.
	   Possible values are: `AND`.
	*/
	LogicalOperator string `json:"logicalOperator,omitempty" yaml:"logicalOperator,omitempty"`
}
