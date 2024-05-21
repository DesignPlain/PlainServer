package types

type Datastream_StreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabaseMysqlTable struct {
	// Table name.
	Table string `json:"table,omitempty" yaml:"table,omitempty"`

	/*
	   MySQL columns in the schema. When unspecified as part of include/exclude objects, includes/excludes everything.
	   Structure is documented below.
	*/
	MysqlColumns []Datastream_StreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabaseMysqlTableMysqlColumn `json:"mysqlColumns,omitempty" yaml:"mysqlColumns,omitempty"`
}
