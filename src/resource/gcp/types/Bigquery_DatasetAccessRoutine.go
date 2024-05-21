package types

type Bigquery_DatasetAccessRoutine struct {
	// The ID of the dataset containing this table.
	DatasetId string `json:"datasetId,omitempty" yaml:"datasetId,omitempty"`

	// The ID of the project containing this table.
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`

	/*
	   The ID of the routine. The ID must contain only letters (a-z,
	   A-Z), numbers (0-9), or underscores (_). The maximum length
	   is 256 characters.
	*/
	RoutineId string `json:"routineId,omitempty" yaml:"routineId,omitempty"`
}
