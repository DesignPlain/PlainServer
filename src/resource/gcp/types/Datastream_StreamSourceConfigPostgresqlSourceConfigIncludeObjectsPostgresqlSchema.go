package types

type Datastream_StreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchema struct {
	// Database name.
	Schema string `json:"schema,omitempty" yaml:"schema,omitempty"`

	/*
	   Tables in the schema.
	   Structure is documented below.
	*/
	PostgresqlTables []Datastream_StreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemaPostgresqlTable `json:"postgresqlTables,omitempty" yaml:"postgresqlTables,omitempty"`
}
