package types

type Lakeformation_DataCellsFilterTableDataRowFilter struct {
	// (Optional) A wildcard that matches all rows.
	AllRowsWildcard Lakeformation_DataCellsFilterTableDataRowFilterAllRowsWildcard `json:"allRowsWildcard,omitempty" yaml:"allRowsWildcard,omitempty"`

	// (Optional) A filter expression.
	FilterExpression string `json:"filterExpression,omitempty" yaml:"filterExpression,omitempty"`
}
