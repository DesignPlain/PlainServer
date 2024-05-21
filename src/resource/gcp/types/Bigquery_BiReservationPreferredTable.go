package types

type Bigquery_BiReservationPreferredTable struct {
	// The ID of the dataset in the above project.
	DatasetId string `json:"datasetId,omitempty" yaml:"datasetId,omitempty"`

	// The assigned project ID of the project.
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`

	// The ID of the table in the above dataset.
	TableId string `json:"tableId,omitempty" yaml:"tableId,omitempty"`
}
