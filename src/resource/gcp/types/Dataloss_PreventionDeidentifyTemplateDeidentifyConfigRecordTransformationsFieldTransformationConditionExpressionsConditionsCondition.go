package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationConditionExpressionsConditionsCondition struct {
	/*
	   Field within the record this condition is evaluated against.
	   Structure is documented below.
	*/
	Field Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationConditionExpressionsConditionsConditionField `json:"field,omitempty" yaml:"field,omitempty"`

	/*
	   Operator used to compare the field or infoType to the value.
	   Possible values are: `EQUAL_TO`, `NOT_EQUAL_TO`, `GREATER_THAN`, `LESS_THAN`, `GREATER_THAN_OR_EQUALS`, `LESS_THAN_OR_EQUALS`, `EXISTS`.
	*/
	Operator string `json:"operator,omitempty" yaml:"operator,omitempty"`

	/*
	   Value to compare against. [Mandatory, except for EXISTS tests.]
	   Structure is documented below.
	*/
	Value Dataloss_PreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationConditionExpressionsConditionsConditionValue `json:"value,omitempty" yaml:"value,omitempty"`
}
