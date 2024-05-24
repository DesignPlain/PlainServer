package types

type Quicksight_DataSetLogicalTableMapDataTransformCastColumnTypeOperation struct {
	// Column name.
	ColumnName string `json:"columnName,omitempty" yaml:"columnName,omitempty"`

	// When casting a column from string to datetime type, you can supply a string in a format supported by Amazon QuickSight to denote the source data format.
	Format string `json:"format,omitempty" yaml:"format,omitempty"`

	// New column data type. Valid values are `STRING`, `INTEGER`, `DECIMAL`, `DATETIME`.
	NewColumnType string `json:"newColumnType,omitempty" yaml:"newColumnType,omitempty"`
}
