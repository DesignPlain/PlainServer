package types

type Bigquery_JobExtractSourceTable struct {
	// The ID of the project containing this table.
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`

	/*
	   The table. Can be specified `{{table_id}}` if `project_id` and `dataset_id` are also set,
	   or of the form `projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}}` if not.
	*/
	TableId string `json:"tableId,omitempty" yaml:"tableId,omitempty"`

	// The ID of the dataset containing this table.
	DatasetId string `json:"datasetId,omitempty" yaml:"datasetId,omitempty"`
}
