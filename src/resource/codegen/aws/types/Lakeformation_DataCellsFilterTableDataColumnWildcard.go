package types

type Lakeformation_DataCellsFilterTableDataColumnWildcard struct {
	// (Optional) Excludes column names. Any column with this name will be excluded.
	ExcludedColumnNames []string `json:"excludedColumnNames,omitempty" yaml:"excludedColumnNames,omitempty"`
}
