package types

type Dataproc_MetastoreServiceMetadataIntegration struct {
	/*
	   The integration config for the Data Catalog service.
	   Structure is documented below.
	*/
	DataCatalogConfig Dataproc_MetastoreServiceMetadataIntegrationDataCatalogConfig `json:"dataCatalogConfig,omitempty" yaml:"dataCatalogConfig,omitempty"`
}
