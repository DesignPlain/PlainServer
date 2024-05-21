package types

type Datastream_StreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchema struct {
	/*
	   Tables in the database.
	   Structure is documented below.
	*/
	OracleTables []Datastream_StreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemaOracleTable `json:"oracleTables,omitempty" yaml:"oracleTables,omitempty"`

	// Schema name.
	Schema string `json:"schema,omitempty" yaml:"schema,omitempty"`
}
