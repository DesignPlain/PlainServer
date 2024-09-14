package types

type Healthcare_FhirStoreStreamConfigBigqueryDestination struct {
	// BigQuery URI to a dataset, up to 2000 characters long, in the format bq://projectId.bqDatasetId
	DatasetUri string `json:"datasetUri,omitempty" yaml:"datasetUri,omitempty"`

	/*
	   The configuration for the exported BigQuery schema.
	   Structure is documented below.
	*/
	SchemaConfig Healthcare_FhirStoreStreamConfigBigqueryDestinationSchemaConfig `json:"schemaConfig,omitempty" yaml:"schemaConfig,omitempty"`
}
