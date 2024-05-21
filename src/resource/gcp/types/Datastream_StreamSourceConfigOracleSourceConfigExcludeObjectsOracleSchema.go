package types

type Datastream_StreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchema struct {
	/*
	   Tables in the database.
	   Structure is documented below.
	*/
	OracleTables []Datastream_StreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemaOracleTable `json:"oracleTables,omitempty" yaml:"oracleTables,omitempty"`

	// Schema name.
	Schema string `json:"schema,omitempty" yaml:"schema,omitempty"`
}
