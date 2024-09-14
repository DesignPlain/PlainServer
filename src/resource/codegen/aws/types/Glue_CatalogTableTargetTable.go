package types

type Glue_CatalogTableTargetTable struct {
	// Name of the catalog database that contains the target table.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// Name of the target table.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Region of the target table.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// ID of the Data Catalog in which the table resides.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`
}
