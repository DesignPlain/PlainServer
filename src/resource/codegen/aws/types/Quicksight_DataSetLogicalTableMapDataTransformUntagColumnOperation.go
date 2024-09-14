package types

type Quicksight_DataSetLogicalTableMapDataTransformUntagColumnOperation struct {
	// Column name.
	ColumnName string `json:"columnName,omitempty" yaml:"columnName,omitempty"`

	// The column tags to remove from this column.
	TagNames []string `json:"tagNames,omitempty" yaml:"tagNames,omitempty"`
}
