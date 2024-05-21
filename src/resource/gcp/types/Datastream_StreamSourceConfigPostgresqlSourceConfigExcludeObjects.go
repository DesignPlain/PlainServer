package types

type Datastream_StreamSourceConfigPostgresqlSourceConfigExcludeObjects struct {
	/*
	   PostgreSQL schemas on the server
	   Structure is documented below.
	*/
	PostgresqlSchemas []Datastream_StreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchema `json:"postgresqlSchemas,omitempty" yaml:"postgresqlSchemas,omitempty"`
}
