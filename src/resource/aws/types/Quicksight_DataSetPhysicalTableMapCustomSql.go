package types

type Quicksight_DataSetPhysicalTableMapCustomSql struct {
	// Column schema from the SQL query result set. See columns.
	Columns []Quicksight_DataSetPhysicalTableMapCustomSqlColumn `json:"columns,omitempty" yaml:"columns,omitempty"`

	// ARN of the data source.
	DataSourceArn string `json:"dataSourceArn,omitempty" yaml:"dataSourceArn,omitempty"`

	// Display name for the SQL query result.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// SQL query.
	SqlQuery string `json:"sqlQuery,omitempty" yaml:"sqlQuery,omitempty"`
}
