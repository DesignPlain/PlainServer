package types

type Datastream_StreamBackfillAllMysqlExcludedObjectsMysqlDatabaseMysqlTable struct {
	/*
	   MySQL columns in the schema. When unspecified as part of include/exclude objects, includes/excludes everything.
	   Structure is documented below.
	*/
	MysqlColumns []Datastream_StreamBackfillAllMysqlExcludedObjectsMysqlDatabaseMysqlTableMysqlColumn `json:"mysqlColumns,omitempty" yaml:"mysqlColumns,omitempty"`

	// Table name.
	Table string `json:"table,omitempty" yaml:"table,omitempty"`
}
