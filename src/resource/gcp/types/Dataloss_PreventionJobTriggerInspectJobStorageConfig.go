package types

type Dataloss_PreventionJobTriggerInspectJobStorageConfig struct {
	/*
	   Configuration of the timespan of the items to include in scanning
	   Structure is documented below.
	*/
	TimespanConfig Dataloss_PreventionJobTriggerInspectJobStorageConfigTimespanConfig `json:"timespanConfig,omitempty" yaml:"timespanConfig,omitempty"`

	/*
	   Options defining BigQuery table and row identifiers.
	   Structure is documented below.
	*/
	BigQueryOptions Dataloss_PreventionJobTriggerInspectJobStorageConfigBigQueryOptions `json:"bigQueryOptions,omitempty" yaml:"bigQueryOptions,omitempty"`

	/*
	   Options defining a file or a set of files within a Google Cloud Storage bucket.
	   Structure is documented below.
	*/
	CloudStorageOptions Dataloss_PreventionJobTriggerInspectJobStorageConfigCloudStorageOptions `json:"cloudStorageOptions,omitempty" yaml:"cloudStorageOptions,omitempty"`

	/*
	   Options defining a data set within Google Cloud Datastore.
	   Structure is documented below.
	*/
	DatastoreOptions Dataloss_PreventionJobTriggerInspectJobStorageConfigDatastoreOptions `json:"datastoreOptions,omitempty" yaml:"datastoreOptions,omitempty"`

	/*
	   Configuration to control jobs where the content being inspected is outside of Google Cloud Platform.
	   Structure is documented below.
	*/
	HybridOptions Dataloss_PreventionJobTriggerInspectJobStorageConfigHybridOptions `json:"hybridOptions,omitempty" yaml:"hybridOptions,omitempty"`
}
