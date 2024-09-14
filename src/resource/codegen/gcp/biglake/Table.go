package biglake

import types "libds/gcp/types"

type Table struct {
	/*
	   Output only. The name of the Table. Format:
	   projects/{project_id_or_number}/locations/{locationId}/catalogs/{catalogId}/databases/{databaseId}/tables/{tableId}


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The database type.
	   Possible values are: `HIVE`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The id of the parent database.
	Database string `json:"database,omitempty" yaml:"database,omitempty"`

	/*
	   Options of a Hive table.
	   Structure is documented below.
	*/
	HiveOptions types.Biglake_TableHiveOptions `json:"hiveOptions,omitempty" yaml:"hiveOptions,omitempty"`
}
