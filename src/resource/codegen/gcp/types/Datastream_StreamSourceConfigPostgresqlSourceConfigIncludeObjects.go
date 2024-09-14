package types

type Datastream_StreamSourceConfigPostgresqlSourceConfigIncludeObjects struct {
	/*
	   PostgreSQL schemas on the server
	   Structure is documented below.
	*/
	PostgresqlSchemas []Datastream_StreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchema `json:"postgresqlSchemas,omitempty" yaml:"postgresqlSchemas,omitempty"`
}
