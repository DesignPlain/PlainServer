package types

type Datastream_StreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabase struct {
	// Database name.
	Database string `json:"database,omitempty" yaml:"database,omitempty"`

	/*
	   Tables in the database.
	   Structure is documented below.
	*/
	MysqlTables []Datastream_StreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabaseMysqlTable `json:"mysqlTables,omitempty" yaml:"mysqlTables,omitempty"`
}
