package types

type Dataplex_DatascanData struct {
	// The Dataplex entity that represents the data source(e.g. BigQuery table) for Datascan.
	Entity string `json:"entity,omitempty" yaml:"entity,omitempty"`

	/*
	   The service-qualified full resource name of the cloud resource for a DataScan job to scan against. The field could be:
	   (Cloud Storage bucket for DataDiscoveryScan)BigQuery table of type "TABLE" for DataProfileScan/DataQualityScan.
	*/
	Resource string `json:"resource,omitempty" yaml:"resource,omitempty"`
}
