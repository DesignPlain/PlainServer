package types

type Bigquery_getDatasetAccessView struct {
	// The ID of the project containing this table.
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`

	/*
	   The ID of the table. The ID must contain only letters (a-z,
	   A-Z), numbers (0-9), or underscores (_). The maximum length
	   is 1,024 characters.
	*/
	TableId string `json:"tableId,omitempty" yaml:"tableId,omitempty"`

	// The dataset ID.
	DatasetId string `json:"datasetId,omitempty" yaml:"datasetId,omitempty"`
}
