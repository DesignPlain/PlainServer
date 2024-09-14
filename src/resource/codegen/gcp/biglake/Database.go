package biglake

import types "libds/gcp/types"

type Database struct {
	// The parent catalog.
	Catalog string `json:"catalog,omitempty" yaml:"catalog,omitempty"`

	/*
	   Options of a Hive database.
	   Structure is documented below.
	*/
	HiveOptions types.Biglake_DatabaseHiveOptions `json:"hiveOptions,omitempty" yaml:"hiveOptions,omitempty"`

	// The name of the database.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The database type.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
