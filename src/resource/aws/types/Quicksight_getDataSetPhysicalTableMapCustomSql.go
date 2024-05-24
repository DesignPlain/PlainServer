package types

type Quicksight_getDataSetPhysicalTableMapCustomSql struct {
	//
	DataSourceArn string `json:"dataSourceArn,omitempty" yaml:"dataSourceArn,omitempty"`

	//
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	SqlQuery string `json:"sqlQuery,omitempty" yaml:"sqlQuery,omitempty"`

	//
	Columns []Quicksight_getDataSetPhysicalTableMapCustomSqlColumn `json:"columns,omitempty" yaml:"columns,omitempty"`
}
