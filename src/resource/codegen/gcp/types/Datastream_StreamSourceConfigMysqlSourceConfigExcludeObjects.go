package types

type Datastream_StreamSourceConfigMysqlSourceConfigExcludeObjects struct {
	/*
	   MySQL databases on the server
	   Structure is documented below.
	*/
	MysqlDatabases []Datastream_StreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabase `json:"mysqlDatabases,omitempty" yaml:"mysqlDatabases,omitempty"`
}
