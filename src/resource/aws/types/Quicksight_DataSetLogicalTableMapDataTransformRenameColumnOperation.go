package types

type Quicksight_DataSetLogicalTableMapDataTransformRenameColumnOperation struct {
	// Column to be renamed.
	ColumnName string `json:"columnName,omitempty" yaml:"columnName,omitempty"`

	// New name for the column.
	NewColumnName string `json:"newColumnName,omitempty" yaml:"newColumnName,omitempty"`
}
