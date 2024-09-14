package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceConfig struct {
	/*
	   Replace each input value with a given value.
	   The `new_value` block must only contain one argument. For example when replacing the contents of a string-type field, only `string_value` should be set.
	   Structure is documented below.
	*/
	NewValue Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceConfigNewValue `json:"newValue,omitempty" yaml:"newValue,omitempty"`
}
