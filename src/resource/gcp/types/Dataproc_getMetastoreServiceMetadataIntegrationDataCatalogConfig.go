package types

type Dataproc_getMetastoreServiceMetadataIntegrationDataCatalogConfig struct {
	// Defines whether the metastore metadata should be synced to Data Catalog. The default value is to disable syncing metastore metadata to Data Catalog.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
