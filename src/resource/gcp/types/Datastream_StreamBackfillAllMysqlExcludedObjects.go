package types

type Datastream_StreamBackfillAllMysqlExcludedObjects struct {
	/*
	   MySQL databases on the server
	   Structure is documented below.
	*/
	MysqlDatabases []Datastream_StreamBackfillAllMysqlExcludedObjectsMysqlDatabase `json:"mysqlDatabases,omitempty" yaml:"mysqlDatabases,omitempty"`
}
