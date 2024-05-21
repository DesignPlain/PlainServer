package types

type Datastream_StreamSourceConfigOracleSourceConfigIncludeObjects struct {
	/*
	   Oracle schemas/databases in the database server
	   Structure is documented below.
	*/
	OracleSchemas []Datastream_StreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchema `json:"oracleSchemas,omitempty" yaml:"oracleSchemas,omitempty"`
}
