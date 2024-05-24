package types

type Quicksight_DataSetLogicalTableMapDataTransformFilterOperation struct {
	// An expression that must evaluate to a Boolean value. Rows for which the expression evaluates to true are kept in the dataset.
	ConditionExpression string `json:"conditionExpression,omitempty" yaml:"conditionExpression,omitempty"`
}
