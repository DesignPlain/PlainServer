package types

type Datastream_StreamSourceConfigMysqlSourceConfigIncludeObjects struct {
	/*
	   MySQL databases on the server
	   Structure is documented below.
	*/
	MysqlDatabases []Datastream_StreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabase `json:"mysqlDatabases,omitempty" yaml:"mysqlDatabases,omitempty"`
}
