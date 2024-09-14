package types

type Bigquery_getDatasetExternalDatasetReference struct {
	// External source that backs this dataset.
	ExternalSource string `json:"externalSource,omitempty" yaml:"externalSource,omitempty"`

	/*
	   The connection id that is used to access the externalSource.
	   Format: projects/{projectId}/locations/{locationId}/connections/{connectionId}
	*/
	Connection string `json:"connection,omitempty" yaml:"connection,omitempty"`
}
