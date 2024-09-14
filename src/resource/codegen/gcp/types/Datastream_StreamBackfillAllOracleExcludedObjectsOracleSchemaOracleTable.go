package types

type Datastream_StreamBackfillAllOracleExcludedObjectsOracleSchemaOracleTable struct {
	/*
	   Oracle columns in the schema. When unspecified as part of include/exclude objects, includes/excludes everything.
	   Structure is documented below.
	*/
	OracleColumns []Datastream_StreamBackfillAllOracleExcludedObjectsOracleSchemaOracleTableOracleColumn `json:"oracleColumns,omitempty" yaml:"oracleColumns,omitempty"`

	// Table name.
	Table string `json:"table,omitempty" yaml:"table,omitempty"`
}
