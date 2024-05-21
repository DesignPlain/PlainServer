package types

type Bigquery_JobQueryDefaultDataset struct {
	/*
	   The dataset. Can be specified `{{dataset_id}}` if `project_id` is also set,
	   or of the form `projects/{{project}}/datasets/{{dataset_id}}` if not.
	*/
	DatasetId string `json:"datasetId,omitempty" yaml:"datasetId,omitempty"`

	// The ID of the project containing this table.
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`
}
