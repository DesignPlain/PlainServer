package types

type Bigquery_JobExtractSourceModel struct {
	// The ID of the dataset containing this model.
	DatasetId string `json:"datasetId,omitempty" yaml:"datasetId,omitempty"`

	/*
	   The ID of the model.

	   - - -
	*/
	ModelId string `json:"modelId,omitempty" yaml:"modelId,omitempty"`

	// The ID of the project containing this model.
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`
}
