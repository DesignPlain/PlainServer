package types

type Datastream_StreamBackfillAllMysqlExcludedObjectsMysqlDatabase struct {
	// Database name.
	Database string `json:"database,omitempty" yaml:"database,omitempty"`

	/*
	   Tables in the database.
	   Structure is documented below.
	*/
	MysqlTables []Datastream_StreamBackfillAllMysqlExcludedObjectsMysqlDatabaseMysqlTable `json:"mysqlTables,omitempty" yaml:"mysqlTables,omitempty"`
}
