package types

type Datastream_StreamBackfillAll struct {
	/*
	   PostgreSQL data source objects to avoid backfilling.
	   Structure is documented below.
	*/
	PostgresqlExcludedObjects Datastream_StreamBackfillAllPostgresqlExcludedObjects `json:"postgresqlExcludedObjects,omitempty" yaml:"postgresqlExcludedObjects,omitempty"`

	/*
	   MySQL data source objects to avoid backfilling.
	   Structure is documented below.
	*/
	MysqlExcludedObjects Datastream_StreamBackfillAllMysqlExcludedObjects `json:"mysqlExcludedObjects,omitempty" yaml:"mysqlExcludedObjects,omitempty"`

	/*
	   PostgreSQL data source objects to avoid backfilling.
	   Structure is documented below.
	*/
	OracleExcludedObjects Datastream_StreamBackfillAllOracleExcludedObjects `json:"oracleExcludedObjects,omitempty" yaml:"oracleExcludedObjects,omitempty"`
}
