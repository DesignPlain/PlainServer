package types

type Dataloss_PreventionJobTriggerInspectJobActionSaveFindingsOutputConfig struct {
	/*
	   Schema used for writing the findings for Inspect jobs. This field is only used for
	   Inspect and must be unspecified for Risk jobs. Columns are derived from the Finding
	   object. If appending to an existing table, any columns from the predefined schema
	   that are missing will be added. No columns in the existing table will be deleted.
	   If unspecified, then all available columns will be used for a new table or an (existing)
	   table with no schema, and no changes will be made to an existing table that has a schema.
	   Only for use with external storage.
	   Possible values are: `BASIC_COLUMNS`, `GCS_COLUMNS`, `DATASTORE_COLUMNS`, `BIG_QUERY_COLUMNS`, `ALL_COLUMNS`.
	*/
	OutputSchema string `json:"outputSchema,omitempty" yaml:"outputSchema,omitempty"`

	/*
	   Information on the location of the target BigQuery Table.
	   Structure is documented below.
	*/
	Table Dataloss_PreventionJobTriggerInspectJobActionSaveFindingsOutputConfigTable `json:"table,omitempty" yaml:"table,omitempty"`
}
