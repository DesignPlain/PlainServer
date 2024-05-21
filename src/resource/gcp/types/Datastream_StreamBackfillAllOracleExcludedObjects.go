package types

type Datastream_StreamBackfillAllOracleExcludedObjects struct {
	/*
	   Oracle schemas/databases in the database server
	   Structure is documented below.
	*/
	OracleSchemas []Datastream_StreamBackfillAllOracleExcludedObjectsOracleSchema `json:"oracleSchemas,omitempty" yaml:"oracleSchemas,omitempty"`
}
