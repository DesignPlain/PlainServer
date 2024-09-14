package types

type Datastream_StreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchema struct {
	/*
	   Tables in the schema.
	   Structure is documented below.
	*/
	PostgresqlTables []Datastream_StreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemaPostgresqlTable `json:"postgresqlTables,omitempty" yaml:"postgresqlTables,omitempty"`

	// Database name.
	Schema string `json:"schema,omitempty" yaml:"schema,omitempty"`
}
