package types

type Quicksight_DataSetLogicalTableMapDataTransformCreateColumnsOperationColumn struct {
	// A unique ID to identify a calculated column. During a dataset update, if the column ID of a calculated column matches that of an existing calculated column, Amazon QuickSight preserves the existing calculated column.
	ColumnId string `json:"columnId,omitempty" yaml:"columnId,omitempty"`

	// Column name.
	ColumnName string `json:"columnName,omitempty" yaml:"columnName,omitempty"`

	// An expression that defines the calculated column.
	Expression string `json:"expression,omitempty" yaml:"expression,omitempty"`
}
