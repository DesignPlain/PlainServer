package types

type Datastream_StreamSourceConfigOracleSourceConfigExcludeObjects struct {
	/*
	   Oracle schemas/databases in the database server
	   Structure is documented below.
	*/
	OracleSchemas []Datastream_StreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchema `json:"oracleSchemas,omitempty" yaml:"oracleSchemas,omitempty"`
}
