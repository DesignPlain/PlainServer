package types

type Datastream_StreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasetsDatasetTemplate struct {
	/*
	   If supplied, every created dataset will have its name prefixed by the provided value.
	   The prefix and name will be separated by an underscore. i.e. _.
	*/
	DatasetIdPrefix string `json:"datasetIdPrefix,omitempty" yaml:"datasetIdPrefix,omitempty"`

	/*
	   Describes the Cloud KMS encryption key that will be used to protect destination BigQuery
	   table. The BigQuery Service Account associated with your project requires access to this
	   encryption key. i.e. projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{cryptoKey}.
	   See https://cloud.google.com/bigquery/docs/customer-managed-encryption for more information.

	   - - -
	*/
	KmsKeyName string `json:"kmsKeyName,omitempty" yaml:"kmsKeyName,omitempty"`

	/*
	   The geographic location where the dataset should reside.
	   See https://cloud.google.com/bigquery/docs/locations for supported locations.
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
}
