package types

type Dataproc_getMetastoreServiceMetadataIntegration struct {
	// The integration config for the Data Catalog service.
	DataCatalogConfigs []Dataproc_getMetastoreServiceMetadataIntegrationDataCatalogConfig `json:"dataCatalogConfigs,omitempty" yaml:"dataCatalogConfigs,omitempty"`
}
