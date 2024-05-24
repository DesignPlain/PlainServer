package types

type Sagemaker_FeatureGroupOfflineStoreConfigDataCatalogConfig struct {
	// The name of the Glue table catalog.
	Catalog string `json:"catalog,omitempty" yaml:"catalog,omitempty"`

	// The name of the Glue table database.
	Database string `json:"database,omitempty" yaml:"database,omitempty"`

	// The name of the Glue table.
	TableName string `json:"tableName,omitempty" yaml:"tableName,omitempty"`
}
