package types

type Dataloss_PreventionJobTriggerInspectJobActionDeidentifyTransformationDetailsStorageConfigTable struct {
	// The ID of the dataset containing this table.
	DatasetId string `json:"datasetId,omitempty" yaml:"datasetId,omitempty"`

	// The ID of the project containing this table.
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`

	/*
	   The ID of the table. The ID must contain only letters (a-z,
	   A-Z), numbers (0-9), or underscores (_). The maximum length
	   is 1,024 characters.
	*/
	TableId string `json:"tableId,omitempty" yaml:"tableId,omitempty"`
}
