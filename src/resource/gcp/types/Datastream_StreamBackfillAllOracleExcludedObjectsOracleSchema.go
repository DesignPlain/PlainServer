package types

type Datastream_StreamBackfillAllOracleExcludedObjectsOracleSchema struct {
	/*
	   Tables in the database.
	   Structure is documented below.
	*/
	OracleTables []Datastream_StreamBackfillAllOracleExcludedObjectsOracleSchemaOracleTable `json:"oracleTables,omitempty" yaml:"oracleTables,omitempty"`

	// Schema name.
	Schema string `json:"schema,omitempty" yaml:"schema,omitempty"`
}
