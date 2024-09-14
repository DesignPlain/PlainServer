package types

type Datastream_StreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemaOracleTable struct {
	/*
	   Oracle columns in the schema. When unspecified as part of include/exclude objects, includes/excludes everything.
	   Structure is documented below.
	*/
	OracleColumns []Datastream_StreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemaOracleTableOracleColumn `json:"oracleColumns,omitempty" yaml:"oracleColumns,omitempty"`

	// Table name.
	Table string `json:"table,omitempty" yaml:"table,omitempty"`
}
