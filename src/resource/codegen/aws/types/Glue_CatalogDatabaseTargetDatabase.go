package types

type Glue_CatalogDatabaseTargetDatabase struct {
	// Region of the target database.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// ID of the Data Catalog in which the database resides.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	// Name of the catalog database.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`
}
