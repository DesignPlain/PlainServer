package types

type Datastream_StreamBackfillAllPostgresqlExcludedObjects struct {
	/*
	   PostgreSQL schemas on the server
	   Structure is documented below.
	*/
	PostgresqlSchemas []Datastream_StreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchema `json:"postgresqlSchemas,omitempty" yaml:"postgresqlSchemas,omitempty"`
}
