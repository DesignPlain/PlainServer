package types

type Glue_getCatalogTableTargetTable struct {
	// ID of the Glue Catalog and database where the table metadata resides. If omitted, this defaults to the current AWS Account ID.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	// Name of the metadata database where the table metadata resides.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// Name of the table.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Region of the target table.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
}
