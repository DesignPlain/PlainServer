package types

type Bigquery_DatasetAccessAuthorizedDatasetDataset struct {
	// The ID of the project containing this table.
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`

	// The ID of the dataset containing this table.
	DatasetId string `json:"datasetId,omitempty" yaml:"datasetId,omitempty"`
}
