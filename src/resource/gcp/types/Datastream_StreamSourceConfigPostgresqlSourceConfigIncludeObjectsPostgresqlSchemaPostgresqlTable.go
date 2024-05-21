package types

type Datastream_StreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemaPostgresqlTable struct {
	/*
	   PostgreSQL columns in the schema. When unspecified as part of include/exclude objects, includes/excludes everything.
	   Structure is documented below.
	*/
	PostgresqlColumns []Datastream_StreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemaPostgresqlTablePostgresqlColumn `json:"postgresqlColumns,omitempty" yaml:"postgresqlColumns,omitempty"`

	// Table name.
	Table string `json:"table,omitempty" yaml:"table,omitempty"`
}
