package types

type Datastream_StreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchema struct {
	/*
	   Tables in the schema.
	   Structure is documented below.
	*/
	PostgresqlTables []Datastream_StreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemaPostgresqlTable `json:"postgresqlTables,omitempty" yaml:"postgresqlTables,omitempty"`

	// Database name.
	Schema string `json:"schema,omitempty" yaml:"schema,omitempty"`
}
